package modules

import (
	"guest/settings"
	"log"
)

type FormatModule struct{}

func NewFormatModule() *FormatModule {
	return new(FormatModule)
}

func (m *FormatModule) GetName() string {
	return "guest/format"
}

func (m *FormatModule) GetExports(data *ModuleData) map[string]interface{} {
	return map[string]interface{}{
		"json": map[string]interface{}{
			"parse": func(input string) interface{} {
				output := make(map[string]interface{})

				err := settings.Parse([]byte(input), &output, settings.FormatJson)
				if err != nil {
					log.Fatalln(err)
				}

				return output
			},
			"stringify": func(input interface{}) string {
				output, err := settings.Stringify(input, settings.FormatJson)
				if err != nil {
					log.Fatalln(err)
				}

				return string(output)
			},
		},
		"yaml": map[string]interface{}{
			"parse": func(input string) interface{} {
				output := make(map[string]interface{})

				err := settings.Parse([]byte(input), &output, settings.FormatYaml)
				if err != nil {
					log.Fatalln(err)
				}

				return output
			},
			"stringify": func(input interface{}) string {
				output, err := settings.Stringify(input, settings.FormatYaml)
				if err != nil {
					log.Fatalln(err)
				}

				return string(output)
			},
		},
	}
}
