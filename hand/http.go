package hand

import (
	"bytes"
	"fmt"
	"guest/body"
	"guest/settings"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type HttpSettings struct {
	SSLVerification  bool   `json:"ssl_verification" title:"SSL certificate verification"`
	FollowRedirects  bool   `json:"follow_redirects" title:"Follow redirects"`
	URLEncode        bool   `json:"url_encode" title:"Url encoding of path and query parameters"`
	MinifyBody       bool   `json:"minify_body" title:"Minification of body (json, xml, etc.)"`
	DisableCookieJar bool   `json:"disable_cookie_jar" title:"Prevent using common cookie jar for request"`
	MaxRedirects     int    `json:"max_redirects" title:"Maximum number of redirects"`
	Proxy            string `json:"proxy" title:"Proxy server URL"`
	Timeout          int    `json:"timeout" title:"Timeout in seconds (0 - forever)"`
}

func (s *HttpSettings) Get(key string) any {
	settingsMap := make(map[string]interface{})
	settings.ReParse(s, &settingsMap)

	return settingsMap[key]
}

func (s *HttpSettings) Set(key string, value any) {
	settingsMap := make(map[string]interface{})
	settings.ReParse(s, &settingsMap)

	settingsMap[key] = value

	settings.ReParse(settingsMap, s)
}

func (s *HttpSettings) MethodsMap() map[string]interface{} {
	return map[string]interface{}{
		"get": s.Get,
		"set": s.Set,
	}
}

type HttpOptions struct {
	Method  string               `json:"method" title:"HTTP method"`
	Url     string               `json:"url" title:"Request URL"`
	Headers settings.KeyValueMap `json:"headers,omitempty" yaml:"headers,omitempty" title:"Request headers"`
	Params  settings.KeyValueMap `json:"params,omitempty" yaml:"params,omitempty" title:"Request parameters"`
	Body    *body.BodyBase       `json:"body,omitempty" yaml:"body,omitempty" title:"Request body"`

	Settings HttpSettings `json:"settings"`

	bodyIO body.Body
}

type HttpHand struct{}

var DefaultHttpOptions = HttpOptions{
	Method: "GET",
	Url:    "",
	Headers: settings.KeyValueMap{
		"Accept": {
			Value: "*/*",
		},
	},
	Params: make(settings.KeyValueMap),

	Settings: HttpSettings{
		SSLVerification:  true,
		FollowRedirects:  false,
		URLEncode:        true,
		MinifyBody:       true,
		DisableCookieJar: false,
		MaxRedirects:     3,
		Proxy:            "",
		Timeout:          5,
	},
}

func NewHttpOptions() *HttpOptions {
	return &HttpOptions{
		Headers: make(settings.KeyValueMap),
		Params:  make(settings.KeyValueMap),

		Settings: HttpSettings{},
	}
}

func (h *HttpHand) Validate(options *HttpOptions) error {
	// if options.Method != "GET" {
	// 	return fmt.Errorf("unsupported method: %v", options.Method)
	// }

	return nil
}

func (h *HttpHand) Run(options *HttpOptions) (*http.Response, error) {
	if err := h.Validate(options); err != nil {
		return nil, err
	}

	resp, err := h.makeRequest(options)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HttpHand) GetExports() map[string]interface{} {
	return map[string]interface{}{
		"get": func(url string) map[string]interface{} {
			return NewHttpRequest("GET", url).MethodsMap()
		},
		"post": func(url string) map[string]interface{} {
			return NewHttpRequest("POST", url).MethodsMap()
		},
		"put": func(url string) map[string]interface{} {
			return NewHttpRequest("PUT", url).MethodsMap()
		},
		"delete": func(url string) map[string]interface{} {
			return NewHttpRequest("DELETE", url).MethodsMap()
		},
	}
}

func (h *HttpHand) makeRequest(options *HttpOptions) (*http.Response, error) {
	var requestBody io.Reader

	requestCount := 0

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			requestCount++

			if options.Settings.FollowRedirects && requestCount < options.Settings.MaxRedirects {
				return nil
			}

			return http.ErrUseLastResponse
		},
		Timeout: time.Duration(options.Settings.Timeout) * time.Second,
	}

	if options.Settings.Proxy != "" {
		proxyUrl, err := url.Parse(options.Settings.Proxy)
		if err != nil {
			return nil, err
		}

		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}

	requestMethod := strings.ToUpper(options.Method)
	requestUrl := strings.Split(options.Url, "?")[0]

	if options.Body != nil && options.Body.GetType() != "" {
		requestBody = bytes.NewReader(options.Body.GetContent())

		if !options.Headers.Has("content-type") {
			options.Headers.Set("content-type", options.Body.ContentType)
		}
	}

	requestQuery := h.makeQuery(options.Params, options.Settings.URLEncode)
	requestPath := requestUrl + requestQuery

	req, err := http.NewRequest(requestMethod, requestPath, requestBody)
	if err != nil {
		return nil, err
	}

	if !options.Headers.Has("accept") {
		options.Headers.Set("accept", "*/*")
	}

	for k, v := range options.Headers {
		req.Header.Set(cases.Title(language.English).String(k), v.Value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HttpHand) makeQuery(params settings.KeyValueMap, encode bool) string {
	paramsCount := len(params)

	if paramsCount == 0 {
		return ""
	}

	if encode {
		requestQuery := url.Values{}
		for k, v := range params {
			if k == "_" {
				continue
			}

			requestQuery.Set(k, v.Value)
		}

		return "?" + requestQuery.Encode()
	}

	parts := make([]string, paramsCount)

	for k, v := range params {
		if k == "_" {
			continue
		}

		parts = append(parts, fmt.Sprintf("%s=%s", k, v.Value))
	}

	return "?" + strings.Join(parts, "&")
}

type HttpRequest struct {
	*HttpOptions
}

func NewHttpRequest(method, url string) *HttpRequest {
	options := NewHttpOptions()
	options.Method = method
	options.Url = url

	return &HttpRequest{
		HttpOptions: options,
	}
}

func (r *HttpRequest) Multipart() map[string]interface{} {
	if r.Body == nil {
		r.Body = new(body.BodyBase)
	}

	r.bodyIO = body.NewMultipartBody(r.Body)

	return r.bodyIO.MethodsMap()
}

func (r *HttpRequest) UrlEncoded() map[string]interface{} {
	if r.Body == nil {
		r.Body = new(body.BodyBase)
	}

	r.bodyIO = body.NewUrlEncodedBody(r.Body)

	return r.bodyIO.MethodsMap()
}

func (r *HttpRequest) Fire() map[string]interface{} {
	h := new(HttpHand)

	if r.Body != nil {
		r.Body.Close()
		r.bodyIO.Close()
	}

	resp, err := h.Run(r.HttpOptions)
	if err != nil {
		panic(err)
	}

	return NewHttpResponse(resp).MethodsMap()
}

func (r *HttpRequest) MethodsMap() map[string]interface{} {
	return map[string]interface{}{
		"settings":   r.Settings.MethodsMap(),
		"headers":    r.Headers.MethodsMap(),
		"params":     r.Params.MethodsMap(),
		"multipart":  r.Multipart,
		"urlencoded": r.UrlEncoded,
		"fire":       r.Fire,
	}
}

type HttpResponse struct {
	*http.Response

	Headers settings.KeyValueMap

	body string
}

func NewHttpResponse(resp *http.Response) *HttpResponse {
	r := &HttpResponse{
		Response: resp,

		Headers: make(settings.KeyValueMap),
	}

	for k, v := range resp.Header {
		r.Headers.Set(k, v[0])
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	r.body = string(body)

	return r
}

func (r *HttpResponse) MethodsMap() map[string]interface{} {
	return map[string]interface{}{
		"status":  r.StatusCode,
		"headers": r.Headers.MethodsMap(),
		"body":    r.body,
	}
}
