package settings

import (
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type KeyValueItem struct {
	Key         string `json:"key,omitempty" yaml:"key,omitempty"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type KeyValueMap map[string]KeyValueItem

func (io KeyValueMap) rebuild() {
	// TODO bad idea
	if _, ok := io["_"]; ok {
		return
	}

	for k, v := range io {
		delete(io, k)

		io[strings.ToLower(k)] = v
	}

	io["_"] = KeyValueItem{}
}

func (io KeyValueMap) Has(key string) bool {
	io.rebuild()

	_, ok := io.Get(key)

	return ok
}

func (io KeyValueMap) Get(key string) (string, bool) {
	io.rebuild()

	lowerKey := strings.ToLower(key)

	v, ok := io[lowerKey]

	return v.Value, ok
}

func (io KeyValueMap) Set(key string, value string) {
	io.rebuild()

	lowerKey := strings.ToLower(key)

	io[lowerKey] = KeyValueItem{
		Value: value,
	}
}

func (io KeyValueMap) Delete(key string) {
	io.rebuild()

	lowerKey := strings.ToLower(key)

	delete(io, lowerKey)
}

func (io KeyValueMap) SortedSlice() []KeyValueItem {
	io.rebuild()

	slice := make([]KeyValueItem, 0)

	for k, v := range io {
		if k == "_" {
			continue
		}

		slice = append(slice, KeyValueItem{
			Key:   cases.Title(language.English).String(k),
			Value: v.Value,
		})
	}

	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i].Key < slice[j].Key
	})

	return slice
}

func (io KeyValueMap) MethodsMap() map[string]interface{} {
	return map[string]interface{}{
		"has":    io.Has,
		"get":    io.Get,
		"set":    io.Set,
		"delete": io.Delete,
	}
}
