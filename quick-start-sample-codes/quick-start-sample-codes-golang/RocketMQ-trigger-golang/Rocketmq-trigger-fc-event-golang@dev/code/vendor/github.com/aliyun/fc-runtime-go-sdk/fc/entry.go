// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package fc

import (
	"context"
	"log"
	"os"
)

// Start takes a handler and talks to an internal fc endpoint to pass requests to the handler. If the
// handler does not match one of the supported types an appropriate error message will be returned to the caller.
// Start blocks, and does not return after being called.
//
// Rules:
//
// 	* handler must be a function
// 	* handler may take between 0 and two arguments.
// 	* if there are two arguments, the first argument must satisfy the "context.Context" interface.
// 	* handler may return between 0 and two arguments.
// 	* if there are two return values, the second argument must be an error.
// 	* if there is one return value it must be an error.
//
// Valid function signatures:
//
// 	func ()
// 	func () error
// 	func (TIn) error
// 	func () (TOut, error)
// 	func (TIn) (TOut, error)
// 	func (context.Context) error
// 	func (context.Context, TIn) error
// 	func (context.Context) (TOut, error)
// 	func (context.Context, TIn) (TOut, error)
//
// Where "TIn" and "TOut" are types compatible with the "encoding/json" standard library.
// See https://golang.org/pkg/encoding/json/#Unmarshal for how deserialization behaves
func Start(handler interface{}) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("start")
	StartWithContext(context.Background(), handler, eventFunction)
}

// StartHttp ...
//
// Valid function signatures:
// func (context.Context, http.ResponseWriter, *http.Request) error
func StartHttp(handler interface{}) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("start http")
	StartWithContext(context.Background(), handler, httpFunction)
}

// RegistryInitializeFunction 注册 initialize 函数，在示例启动时执行且只执行一次
func RegisterInitializerFunction(handler interface{}) {
	registerLifeCycleHandler(handler, initializerFunction)
}

// RegisterPreFreezeFunction ...
func RegisterPreFreezeFunction(handler interface{}) {
	registerLifeCycleHandler(handler, preFreezeFunction)
}

// RegisterPreStopFunction ...
func RegisterPreStopFunction(handler interface{}) {
	registerLifeCycleHandler(handler, preStopFunction)
}

// StartWithContext is the same as Start except sets the base context for the function.
func StartWithContext(ctx context.Context, handler interface{}, funcType functionType) {
	StartHandlerWithContext(ctx, handler, funcType)
}

// StartHandler takes in a Handler wrapper interface which can be implemented either by a
// custom function or a struct.
//
// Handler implementation requires a single "Invoke()" function:
//
//  func Invoke(context.Context, []byte) ([]byte, error)
func StartHandler(handler Handler) {
	StartHandlerWithContext(context.Background(), handler, eventFunction)
}

func registerLifeCycleHandler(handler interface{}, funcType functionType) {
	lifeCycleHandlers = append(lifeCycleHandlers, handlerWrapper{handler, funcType})
}

type startFunction struct {
	env string
	f   func(ctx context.Context, envValue string, handler handlerWrapper, lifeHandler []handlerWrapper) error
}

var (
	runtimeAPIStartFunction = &startFunction{
		env: "FC_RUNTIME_API",
		f:   startRuntimeAPILoop,
	}

	lifeCycleHandlers = []handlerWrapper{}

	// This allows end to end testing of the Start functions
	logFatalf = log.Fatalf
	logPrintf = log.Printf

	enableInvokePanicLog = true
)

// StartHandlerWithContext is the same as StartHandler except sets the base context for the function.
//
// Handler implementation requires a single "Invoke()" function:
//
//  func Invoke(context.Context, []byte) ([]byte, error)
func StartHandlerWithContext(ctx context.Context, handler interface{}, funcType functionType) {
	startFunction := runtimeAPIStartFunction
	config := os.Getenv(startFunction.env)
	if config != "" {
		// in normal operation, the start function never returns
		// if it does, exit!, this triggers a restart of the lambda function
		err := startFunction.f(ctx, config, handlerWrapper{handler, funcType}, lifeCycleHandlers)
		logFatalf("%v", err)
	}
	logFatalf("expected ali FC environment variables [%s] are not defined", startFunction.env)
}
