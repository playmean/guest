package settings

import (
	"fmt"
)

type MapIO map[string]string

func (io MapIO) String() string {
	output := ""

	for k, v := range io {
		output += fmt.Sprintf("%s=%s", k, v)
	}

	return output
}

func (io MapIO) Has(key string) bool {
	_, ok := io.Get(key)

	return ok
}

func (io MapIO) Get(key string) (string, bool) {
	v, ok := io[key]

	return v, ok
}

func (io MapIO) Set(key string, value string) {
	io[key] = value
}

func (io MapIO) Delete(key string) {
	delete(io, key)
}

func (io MapIO) MethodsMap() map[string]interface{} {
	return map[string]interface{}{
		"has":    io.Has,
		"get":    io.Get,
		"set":    io.Set,
		"delete": io.Delete,
	}
}
