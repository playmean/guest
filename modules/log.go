package modules

import "log"

type LogModule struct{}

func NewLogModule() *LogModule {
	return new(LogModule)
}

func (m *LogModule) GetName() string {
	return "guest/log"
}

func (m *LogModule) GetExports(data *ModuleData) map[string]interface{} {
	return map[string]interface{}{
		"info": log.Println,
		"panic": func(message string) {
			panic(message)
		},
	}
}
