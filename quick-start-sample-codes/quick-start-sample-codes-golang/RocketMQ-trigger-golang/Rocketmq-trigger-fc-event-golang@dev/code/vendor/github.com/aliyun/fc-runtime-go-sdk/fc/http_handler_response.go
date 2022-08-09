// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package fc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// fcResponse implements http.ResponseWriter.
//
// The response body not support stream write.
type fcResponse struct {
	req         *http.Request
	status      int
	wroteHeader bool
	header      http.Header
	body        bytes.Buffer

	contentLength int64 // explicitly-declared Content-Length; or -1
}

func newFcResponse(req *http.Request) *fcResponse {
	return &fcResponse{
		req:           req,
		contentLength: -1,
		header:        http.Header{},
		body:          bytes.Buffer{},
	}
}

func (r *fcResponse) Header() http.Header {
	return r.header
}

func (r *fcResponse) Write(p []byte) (n int, err error) {
	if !r.wroteHeader {
		r.WriteHeader(http.StatusOK)
	}
	return r.body.Write(p)
}

func (r *fcResponse) WriteHeader(statusCode int) {
	if r.wroteHeader {
		return
	}
	r.wroteHeader = true
	r.status = statusCode
	if r.header.Get("Date") == "" {
		r.header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	}
}

func (r *fcResponse) responseParams() (string, error) {
	respParams := map[string]interface{}{
		"status":     r.status,
		"headersMap": r.header,
	}
	respParamsStr, err := json.Marshal(respParams)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(respParamsStr)
	return encoded, nil
}

// Deprecated: Use Body and HttpParam instead
func (r *fcResponse) Payload() ([]byte, error) {
	respHeaders := map[string]string{}
	for key, values := range r.header {
		respHeaders[key] = strings.Join(values, ", ")
	}
	encodedHttpParams, err := r.responseParams()
	if err != nil {
		return []byte{}, err
	}
	respHeaders[headerHttpParams] = encodedHttpParams

	// json.Marshal encoded []byte as a base64-encoded string
	// so `base64.StdEncoding.EncodeToString(r.body.Bytes())` is not necessary
	// see https://pkg.go.dev/encoding/json#Marshal
	resp := map[string]interface{}{
		"headers":         respHeaders,
		"body":            r.body.Bytes(),
		"isBase64Encoded": true,
	}
	return json.Marshal(resp)
}

func (r *fcResponse) Body() []byte {
	body := r.body.Bytes()
	r.body.Reset()
	return body
}

func (r *fcResponse) HttpParam() (string, error)  {
	encodedHttpParams, err := r.responseParams()
	if err != nil {
		return "", err
	}
	return encodedHttpParams, nil
}
