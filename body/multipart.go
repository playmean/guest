package body

import (
	"bytes"
	"mime/multipart"
)

type MultipartBody struct {
	*BodyBase

	buffer *bytes.Buffer
	writer *multipart.Writer
}

const MultipartContentType = "multipart/form-data"

func NewMultipartBody(base *BodyBase) *MultipartBody {
	b := MultipartBody{}
	b.BodyBase = base
	b.ContentType = MultipartContentType
	b.Content = make([]byte, 0)

	b.buffer = new(bytes.Buffer)
	b.writer = multipart.NewWriter(b.buffer)

	return &b
}

func (b *MultipartBody) Add(key, value string) error {
	err := b.writer.WriteField(key, value)
	if err != nil {
		return err
	}

	return nil
}

func (b *MultipartBody) Close() error {
	b.Content = b.buffer.Bytes()

	return b.writer.Close()
}

func (b *MultipartBody) MethodsMap() map[string]interface{} {
	return map[string]interface{}{
		"add": b.Add,
	}
}
