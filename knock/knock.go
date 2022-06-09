package knock

import (
	"errors"
	"fmt"
	"guest/core"
	"guest/hand"
	"guest/modules"
	"guest/settings"
	"guest/storage"
	"io/ioutil"
	"strings"

	"github.com/imdario/mergo"
)

type Knock interface {
	Validate() error
	Run() (*Result, error)
	RunScript(externalScripts map[string]string, vars map[string]string, scriptType string) error
	PatchOptions(handOptions interface{}) error
	ApplyVariables(vars map[string]string) error
	GetType() string
	GetExports() map[string]interface{}
	GetHandExports() map[string]interface{}
}

type KnockBase struct {
	Description string `json:"description"`

	Type    string            `json:"type"`
	Scripts map[string]string `json:"scripts,omitempty" yaml:"scripts,omitempty"`

	virtualFs storage.Storage
	custom    Knock
}

// TODO move folders to knock

func FromFile(path string, vfs storage.Storage) (Knock, error) {
	f, err := vfs.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return FromData(vfs, data, settings.FormatJson)
}

func FromData(vfs storage.Storage, data []byte, format settings.SettingsFormat) (Knock, error) {
	base := new(KnockBase)

	err := settings.Parse(data, base, format)
	if err != nil {
		return nil, err
	}

	switch base.Type {
	case "http":
		knock := NewHttpKnock(&hand.HttpHand{}, hand.NewHttpOptions())
		knock.virtualFs = vfs

		err := settings.Parse(data, knock, format)
		if err != nil {
			return nil, err
		}

		return knock, nil
	default:
		return nil, fmt.Errorf("unknown knock type: %v", base.Type)
	}
}

func (k *KnockBase) Validate() error {
	return errors.New("knock.Validate not implemented")
}

func (k *KnockBase) Run() (*Result, error) {
	return nil, errors.New("knock.Run not implemented")
}

func (k *KnockBase) RunScript(
	externalScripts map[string]string,
	vars map[string]string,
	scriptType string,
) error {
	scriptPath, ok := k.Scripts[scriptType]
	if !ok {
		scriptPath, ok = externalScripts[scriptType]
		if !ok {
			return nil
		}
	}

	// TODO 10.06.2022 load scripts from URL and inline
	if !storage.Exists(scriptPath, k.virtualFs) {
		return fmt.Errorf("script with type '%s' not found: %s", scriptType, scriptPath)
	}

	f, err := k.virtualFs.Open(scriptPath)
	if err != nil {
		return err
	}

	defer f.Close()

	source, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	knockData := make(map[string]interface{})

	err = settings.ReParse(k.custom, &knockData)
	if err != nil {
		return err
	}

	exports := k.custom.GetExports()
	exports["hand"] = k.custom.GetHandExports()

	knockData, err = executeCore(source, &modules.ModuleData{
		KnockType:    k.Type,
		KnockData:    knockData,
		KnockExports: exports,
		Variables:    vars,
	})
	if err != nil {
		return err
	}

	err = settings.ReParse(knockData, k.custom)
	if err != nil {
		return err
	}

	return nil
}

func (k *KnockBase) PatchOptions(handOptions interface{}) error {
	patch := make(map[string]interface{})
	patch["options"] = handOptions

	mapped := make(map[string]interface{})

	err := settings.ReParse(k.custom, &mapped)
	if err != nil {
		return err
	}

	err = mergo.Merge(&mapped, patch)
	if err != nil {
		return err
	}

	err = settings.ReParse(mapped, &k.custom)
	if err != nil {
		return err
	}

	return nil
}

func (k *KnockBase) ApplyVariables(vars map[string]string) error {
	data, err := settings.Stringify(k.custom, settings.FormatJson)
	if err != nil {
		return err
	}

	dataString := string(data)

	for k, v := range vars {
		dataString = strings.ReplaceAll(dataString, "$"+k, v)
	}

	err = settings.Parse([]byte(dataString), k.custom, settings.FormatJson)
	if err != nil {
		return err
	}

	return nil
}

func (k *KnockBase) GetType() string {
	return k.Type
}

func executeCore(
	source []byte,
	data *modules.ModuleData,
) (map[string]interface{}, error) {
	c := core.NewCore(data)

	err := c.Init()
	if err != nil {
		return nil, err
	}

	defer c.Destroy()

	err = c.Execute(string(source))
	if err != nil {
		return nil, err
	}

	return c.Data.KnockData, nil
}
