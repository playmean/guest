package modules

type KnockModule struct{}

func NewKnockModule() *KnockModule {
	return new(KnockModule)
}

func (m *KnockModule) GetName() string {
	return "guest/knock"
}

func (m *KnockModule) GetExports(data *ModuleData) map[string]interface{} {
	return data.KnockExports
}
