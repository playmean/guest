package settings

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// TODO detect settings type from file

type SettingsFormat string

const (
	FormatJson SettingsFormat = "json"
	FormatYaml SettingsFormat = "yaml"
)

func Stringify(input interface{}, format SettingsFormat) ([]byte, error) {
	var out []byte
	var err error

	switch format {
	case FormatJson:
		out, err = json.MarshalIndent(input, "", "    ")
	case FormatYaml:
		out, err = yaml.Marshal(input)
	default:
		err = fmt.Errorf("unknown stringify format: %s", format)
	}

	return out, err
}

func Parse(input []byte, output interface{}, format SettingsFormat) error {
	var err error

	switch format {
	case FormatJson:
		err = json.Unmarshal(input, output)
	case FormatYaml:
		err = yaml.Unmarshal(input, output)
	default:
		err = fmt.Errorf("unknown parse format: %s", format)
	}

	return err
}

func ReParse(input interface{}, output interface{}) error {
	out, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(out, output)
}
