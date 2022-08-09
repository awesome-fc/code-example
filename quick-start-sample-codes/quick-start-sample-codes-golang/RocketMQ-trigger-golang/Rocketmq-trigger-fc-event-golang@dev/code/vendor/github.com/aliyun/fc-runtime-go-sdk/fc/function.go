// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package fc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc/messages"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

type functionType uint8

const (
	// function type, do not modify!!
	handleFunction      functionType = 0
	initializerFunction functionType = 1
	preStopFunction     functionType = 2
	preFreezeFunction   functionType = 3

	// base function type
	eventFunction functionType = 101
	httpFunction  functionType = 102
)

// Function struct which wrap the Handler
type Function struct {
	ctx         context.Context
	funcType    functionType
	handler     Handler
	httpHandler HttpHandler

	// life cycle handler
	initializeHandler LifeCycleHandler
	preFreezeHandler  LifeCycleHandler
	preStopHandler    LifeCycleHandler
}

// NewFunction which creates a Function with a given Handler
func NewFunction(handler interface{}, funcType functionType) *Function {
	f := &Function{
		funcType: funcType,
	}
	if f.funcType == eventFunction {
		f.handler = NewHandler(handler)
	} else {
		f.httpHandler = NewHttpHandler(handler)
	}
	return f
}

// RegistryLifeCycleHandler ...
func (fn *Function) RegistryLifeCycleHandler(lifeCycleHandlers []handlerWrapper) {
	for _, item := range lifeCycleHandlers {
		switch item.funcType {
		case initializerFunction:
			fn.initializeHandler = NewLifeCycleHandler(item.handler)
		case preFreezeFunction:
			fn.preFreezeHandler = NewLifeCycleHandler(item.handler)
		case preStopFunction:
			fn.preStopHandler = NewLifeCycleHandler(item.handler)
		default:
			// TODO ...
		}
	}
}

// Ping method which given a PingRequest and a PingResponse parses the PingResponse
func (fn *Function) Ping(req *messages.PingRequest, response *messages.PingResponse) error {
	*response = messages.PingResponse{}
	return nil
}

// Invoke method try to perform a command given an InvokeRequest and an InvokeResponse
func (fn *Function) Invoke(req *messages.InvokeRequest, response *messages.InvokeResponse, invokeFuncType functionType) (err error) {
	defer func() {
		if e := recover(); e != nil {
			response.Error = fcPanicResponse(e)
			fn.printPanicLog(req.RequestId, response.Error.ToJson())
			fn.printEndLog(invokeFuncType, req.RequestId, false)
			err = fmt.Errorf("%v", e)
		} else {
			fn.printEndLog(invokeFuncType, req.RequestId, true)
		}
	}()

	deadline := time.Unix(req.Deadline.Seconds, req.Deadline.Nanos).UTC()
	invokeContext, cancel := context.WithDeadline(fn.context(), deadline)
	defer cancel()
	lc := &req.Context
	lc.RequestID = req.RequestId
	invokeContext = fccontext.NewContext(invokeContext, lc)
	fn.printStartLog(invokeFuncType, req.RequestId)
	if invokeFuncType == initializerFunction {
		if fn.initializeHandler == nil {
			fn.initializeHandler = errorLifeCycleHandler(errors.New("no initializer handler registered"))
		}

		fn.initializeHandler.Invoke(invokeContext)
		return nil
	}

	if invokeFuncType == preFreezeFunction {
		if fn.preFreezeHandler == nil {
			fn.preFreezeHandler = errorLifeCycleHandler(errors.New("no prefreeze handler registered"))
		}
		fn.preFreezeHandler.Invoke(invokeContext)
		return nil
	}

	if invokeFuncType == preStopFunction {
		if fn.preStopHandler == nil {
			fn.preStopHandler = errorLifeCycleHandler(errors.New("no prestop handler registered"))
		}
		fn.preStopHandler.Invoke(invokeContext)
		response.Payload = []byte{}
		return nil
	}

	if fn.funcType == eventFunction {
		return fn.invokeEventFunc(invokeContext, req.Payload, response)
	}
	return fn.invokeHttpFunc(invokeContext, req.HttpParams, req.Payload, response)
}

func (fn *Function) invokeHttpFunc(invokeContext context.Context, httpParams *string,
	reqPayload []byte, response *messages.InvokeResponse) error {
	if httpParams == nil {
		handler := errorHttpHandler(fmt.Errorf("no httpParams found in request"))
		err := handler(invokeContext, newFcResponse(&http.Request{}), &http.Request{})
		response.Error = fcErrorResponse(err)
		return nil
	}
	req, err := genHttpRequest(*httpParams, reqPayload)
	if err != nil {
		response.Error = fcErrorResponse(err)
		return nil
	}
	resp := newFcResponse(req)
	err = fn.httpHandler.Invoke(invokeContext, resp, req)
	if err != nil {
		response.Error = fcErrorResponse(err)
		return nil
	}
	response.Payload = resp.Body()
	response.HttpParam, err = resp.HttpParam()
	if err != nil {
		response.Error = fcErrorResponse(err)
		return nil
	}
	return nil
}

func (fn *Function) invokeEventFunc(invokeContext context.Context, reqPayload []byte, response *messages.InvokeResponse) error {
	respPayload, err := fn.handler.Invoke(invokeContext, reqPayload)
	if err != nil {
		response.Error = fcErrorResponse(err)
		return nil
	}
	response.Payload = respPayload
	return nil
}

// context returns the base context used for the fn.
func (fn *Function) context() context.Context {
	if fn.ctx == nil {
		return context.Background()
	}

	return fn.ctx
}

// withContext returns a shallow copy of Function with its context changed
// to the provided ctx. If the provided ctx is non-nil a Background context is set.
func (fn *Function) withContext(ctx context.Context) *Function {
	if ctx == nil {
		ctx = context.Background()
	}

	fn2 := new(Function)
	*fn2 = *fn

	fn2.ctx = ctx

	return fn2
}

func (fn *Function) printPanicLog(requestId, errorMessage string) {
	if !enableInvokePanicLog {
		return
	}
	log.Printf(" %s [ERROR] %s\n", requestId, errorMessage)
}

func (fn *Function) printEndLog(funcType functionType, requestId string, isHandled bool) {
	suffix := ""
	if !isHandled {
		suffix = ", Error: Unhandled function error"
	}

	switch funcType {
	case initializerFunction:
		fmt.Printf("FC Initialize End RequestId: %s%s\n", requestId, suffix)
	case preStopFunction:
		fmt.Printf("FC PreStop End RequestId: %s%s\n", requestId, suffix)
	case preFreezeFunction:
		fmt.Printf("FC PreFreeze End RequestId: %s%s\n", requestId, suffix)
	default:
		fmt.Printf("FC Invoke End RequestId: %s%s\n", requestId, suffix)
	}
}

func (fn *Function) printStartLog(funcType functionType, requestId string) {
	switch funcType {
	case initializerFunction:
		fmt.Printf("FC Initialize Start RequestId: %s\n", requestId)
	case preStopFunction:
		fmt.Printf("FC PreStop Start RequestId: %s\n", requestId)
	case preFreezeFunction:
		fmt.Printf("FC PreFreeze Start RequestId: %s\n", requestId)
	default:
		fmt.Printf("FC Invoke Start RequestId: %s\n", requestId)
	}
}
