package body

import "guest/settings"

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
	stringContent, ok := b.Content.(string)
	if ok {
		return []byte(stringContent)
	}

	bytesContent, ok := b.Content.([]byte)
	if ok {
		return bytesContent
	}

	bytesContent, err := settings.Stringify(b.Content, settings.FormatJson)
	if err == nil {
		return bytesContent
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
