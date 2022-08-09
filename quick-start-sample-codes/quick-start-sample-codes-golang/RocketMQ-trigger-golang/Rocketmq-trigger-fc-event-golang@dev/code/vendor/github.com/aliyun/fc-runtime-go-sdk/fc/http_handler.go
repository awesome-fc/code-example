// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.
// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package fc

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type HttpHandler interface {
	Invoke(ctx context.Context, w http.ResponseWriter, req *http.Request) error
}

// fcHandler is the generic function type
type fcHttpHandler func(context.Context, http.ResponseWriter, *http.Request) error

// Invoke calls the handler, and serializes the response.
// If the underlying handler returned an error, or an error occurs during serialization, error is returned.
func (handler fcHttpHandler) Invoke(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	return handler(ctx, w, req)
}

func errorHttpHandler(e error) fcHttpHandler {
	return func(ctx context.Context, rsp http.ResponseWriter, req *http.Request) error {
		return e
	}
}

func validateHttpArguments(handler reflect.Type) error {
	if handler.NumIn() != 3 {
		return fmt.Errorf("http handler should takes three arguments, but handler takes %d", handler.NumIn())
	}
	contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
	argumentType := handler.In(0)
	if !argumentType.Implements(contextType) {
		return fmt.Errorf("handler takes three arguments, but the first is not Context. got %s",
			argumentType.Kind())
	}
	responseWriterType := reflect.TypeOf((*http.ResponseWriter)(nil)).Elem()
	argumentType = handler.In(1)
	if !argumentType.Implements(responseWriterType) {
		return fmt.Errorf("handler takes three arguments, but the second is not http.ResponseWriter. got %s",
			argumentType.Kind())
	}
	return nil
}

func validateHttpReturns(handler reflect.Type) error {
	errorType := reflect.TypeOf((*error)(nil)).Elem()

	switch n := handler.NumOut(); {
	case n > 1:
		return fmt.Errorf("handler may not return more than one values")
	case n == 1:
		if !handler.Out(0).Implements(errorType) {
			return fmt.Errorf("handler returns a single value, but it does not implement error")
		}
	}

	return nil
}

// NewHttpHandler creates a http fc handler from the given handler function. The
// returned Handler performs JSON serialization and deserialization, and
// delegates to the input handler function.  The handler function parameter must
// satisfy the rules documented by Start.  If handlerFunc is not a valid
// handler, the returned Handler simply reports the validation error.
func NewHttpHandler(handlerFunc interface{}) HttpHandler {
	if handlerFunc == nil {
		return errorHttpHandler(fmt.Errorf("handler is nil"))
	}
	handler := reflect.ValueOf(handlerFunc)
	handlerType := reflect.TypeOf(handlerFunc)
	if handlerType.Kind() != reflect.Func {
		return errorHttpHandler(fmt.Errorf("handler kind %s is not %s", handlerType.Kind(), reflect.Func))
	}

	err := validateHttpArguments(handlerType)
	if err != nil {
		return errorHttpHandler(err)
	}

	if err := validateHttpReturns(handlerType); err != nil {
		return errorHttpHandler(err)
	}

	return fcHttpHandler(func(ctx context.Context, rsp http.ResponseWriter, req *http.Request) error {
		// construct arguments
		args := []reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(rsp),
			reflect.ValueOf(req),
		}

		response := handler.Call(args)

		// convert return values into ( error)
		var err error
		if len(response) > 0 {
			if errVal, ok := response[len(response)-1].Interface().(error); ok {
				err = errVal
			}
		}

		return err
	})
}

type HTTPParams struct {
	Path       string              `json:"path"`
	Method     string              `json:"method"`
	RequestURI string              `json:"requestURI"`
	ClientIP   string              `json:"clientIP"`
	Host       string              `json:"host"`
	QueriesMap map[string][]string `json:"queriesMap"`
	HeadersMap map[string][]string `json:"headersMap"`
}

func genHttpRequest(httpParams string, payload []byte) (*http.Request, error) {
	data, err := base64.StdEncoding.DecodeString(httpParams)
	if err != nil {
		return nil, err
	}
	var params HTTPParams
	if err = json.Unmarshal(data, &params); err != nil {
		return nil, err
	}
	// generate req.RequestURI, req.URL.{Path, RawQuery}
	req, err := http.NewRequest(params.Method, params.RequestURI, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.URL.Path = params.Path
	req.RemoteAddr = params.ClientIP
	req.Host = params.ClientIP
	req.Header = params.HeadersMap
	req.RequestURI = params.RequestURI

	return req, nil
}
