package knock

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/playmean/guest/body"
	"github.com/playmean/guest/hand"
	"github.com/playmean/guest/settings"
)

type HttpKnock struct {
	KnockBase

	Hand    *hand.HttpHand    `json:"-" yaml:"-"`
	Options *hand.HttpOptions `json:"options"`
}

func NewHttpKnock(h *hand.HttpHand, options *hand.HttpOptions) *HttpKnock {
	k := new(HttpKnock)
	k.Type = "http"
	k.Hand = h
	k.Options = options
	k.custom = k

	return k
}

func (k *HttpKnock) Validate() error {
	return k.Hand.Validate(k.Options)
}

func (k *HttpKnock) Run() (*Result, error) {
	handStart := time.Now()

	resp, err := k.Hand.Run(k.Options)
	if err != nil {
		return nil, err
	}

	handTime := time.Since(handStart)

	bodyBuffer, _ := ioutil.ReadAll(resp.Body)
	headers := make(settings.KeyValueMap)
	statusFlags := make([]StatusFlag, 0)

	switch resp.StatusCode / 100 {
	case 1, 2:
		statusFlags = append(statusFlags, StatusFlagGood)
	case 3:
		statusFlags = append(statusFlags, StatusFlagWarning)
	default:
		statusFlags = append(statusFlags, StatusFlagBad)
	}

	contentLen := int(resp.ContentLength)
	if contentLen < 0 {
		contentLen = len(bodyBuffer)
	}

	for k := range resp.Header {
		headers.Set(k, resp.Header.Get(k))
	}

	entities := make([]ResultEntity, 0)

	if len(bodyBuffer) > 0 {
		entities = append(entities, ResultEntity{
			Title: "Body",
			Entity: body.ResolveBody(&body.BodyBase{
				ContentType: resp.Header.Get("Content-Type"),
				Content:     bodyBuffer,
			}),
		})
	}

	entities = append(entities, ResultEntity{
		Title:  "Headers",
		Entity: headers,
	})

	result := &Result{
		Entities: entities,
		Statuses: []ResultStatus{
			{
				Title: "Status",
				Value: resp.Status,
				Flags: statusFlags,
			},
			{
				Title: "Size",
				Value: fmt.Sprintf("%d bytes", contentLen),
			},
			{
				Title: "Time",
				Value: fmt.Sprintf("%d ms", handTime.Milliseconds()),
			},
		},
	}

	return result, nil
}

func (k *HttpKnock) GetExports() map[string]interface{} {
	return map[string]interface{}{
		"headers": k.Options.Headers.MethodsMap(),
		"params":  k.Options.Params.MethodsMap(),
	}
}

func (k *HttpKnock) GetHandExports() map[string]interface{} {
	return k.Hand.GetExports()
}
