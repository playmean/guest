package modules

import "github.com/playmean/guest/settings"

type ModuleData struct {
	KnockType    string
	KnockData    map[string]interface{}
	KnockExports map[string]interface{}
	Variables    settings.MapIO
}

type Module interface {
	GetName() string
	GetExports(*ModuleData) map[string]interface{}
}

var Internal = []Module{
	NewGuestModule(),
	NewKnockModule(),
	NewLogModule(),
	NewFormatModule(),
	NewVariablesModule(),
}
