package modules

type VariablesModule struct{}

func NewVariablesModule() *VariablesModule {
	return new(VariablesModule)
}

func (m *VariablesModule) GetName() string {
	return "guest/variables"
}

func (m *VariablesModule) GetExports(data *ModuleData) map[string]interface{} {
	return data.Variables.MethodsMap()
}
