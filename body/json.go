package body

import (
	"github.com/playmean/guest/settings"
)

type JsonBody struct {
	*BodyBase
}

const JsonContentType = "application/json"

func NewJsonBody(base *BodyBase) *JsonBody {
	b := JsonBody{}
	b.BodyBase = base

	return &b
}

func (b *JsonBody) String() string {
	bytesContent, ok := b.Content.([]byte)
	if ok {
		var dummy interface{}

		err := settings.Parse(bytesContent, &dummy, settings.FormatJson)
		if err != nil {
			panic(err)
		}

		bytesContent, err = settings.Stringify(dummy, settings.FormatJson)
		if err != nil {
			panic(err)
		}

		return string(bytesContent)
	}

	return ""
}

func (b *JsonBody) Close() error {
	return nil
}
