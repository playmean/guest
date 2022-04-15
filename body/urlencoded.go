package body

import (
	"net/url"
)

type UrlEncodedBody struct {
	*BodyBase

	values url.Values
}

const UrlEncodedContentType = "application/x-www-form-urlencoded"

func NewUrlEncodedBody(base *BodyBase) *UrlEncodedBody {
	b := UrlEncodedBody{}
	b.BodyBase = base
	b.ContentType = UrlEncodedContentType
	b.Content = make([]byte, 0)

	b.values = url.Values{}

	return &b
}

func (b *UrlEncodedBody) Add(key, value string) error {
	b.values.Add(key, value)

	return nil
}

func (b *UrlEncodedBody) Close() error {
	b.Content = []byte(b.values.Encode())

	return nil
}

func (b *UrlEncodedBody) MethodsMap() map[string]interface{} {
	return map[string]interface{}{
		"add": b.Add,
	}
}
