package modules

type GuestModule struct{}

func NewGuestModule() *GuestModule {
	return new(GuestModule)
}

func (m *GuestModule) GetName() string {
	return "guest"
}

func (m *GuestModule) GetExports(data *ModuleData) map[string]interface{} {
	return map[string]interface{}{
		"version": "dev",
	}
}
