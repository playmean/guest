package hand

type Hand interface {
	Validate(interface{}) error
	Run(interface{}) (interface{}, error)
	GetExports() map[string]interface{}
}
