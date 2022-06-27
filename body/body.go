package body

import (
	"github.com/playmean/guest/settings"
)

type Body interface {
	GetType() string
	GetContent() []byte
	String() string
	Close() error
	MethodsMap() map[string]interface{}
}

type BodyBase struct {
	ContentType string      `json:"content_type"`
	Content     interface{} `json:"content"`
}

func (b *BodyBase) GetType() string {
	return b.ContentType
}

func (b *BodyBase) GetContent() []byte {
	switch raw := b.Content.(type) {
	case string:
		return []byte(raw)
	case []byte:
		return raw
	default:
		if bytesContent, err := settings.Stringify(b.Content, settings.FormatJson); err == nil {
			return bytesContent
		}
	}

	return nil
}

func (b *BodyBase) String() string {
	return string(b.GetContent())
}

func (b *BodyBase) Close() error {
	return nil
}

func (b *BodyBase) MethodsMap() map[string]interface{} {
	return map[string]interface{}{}
}
