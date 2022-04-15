package body

import (
	"strings"
)

func ResolveBody(base *BodyBase) Body {
	// TODO remove (?)
	if strings.Contains(base.ContentType, "json") {
		return NewJsonBody(base)
	}

	switch base.ContentType {
	case MultipartContentType:
		return NewMultipartBody(base)
	case UrlEncodedContentType:
		return NewUrlEncodedBody(base)
	case JsonContentType:
		return NewJsonBody(base)
	default:
		return base
	}
}
