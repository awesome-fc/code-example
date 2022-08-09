// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package fc

import (
	"context"
	"fmt"
	"reflect"
)

// include initialize handler, pre-freeze handler and pre-stop handler
type LifeCycleHandler interface {
	Invoke(ctx context.Context)
}

// fcLifeCycleHandler is the generic function type
type fcLifeCycleHandler func(context.Context)

// Invoke calls the handler, and serializes the response.
// If the underlying handler returned an error, or an error occurs during serialization, error is returned.
func (handler fcLifeCycleHandler) Invoke(ctx context.Context) {
	handler(ctx)
}

func errorLifeCycleHandler(e error) fcLifeCycleHandler {
	return func(ctx context.Context) {
		panic(e)
	}
}

func validateLifeCycleArguments(handler reflect.Type) error {
	if handler.NumIn() != 1 {
		return fmt.Errorf("life cycle handler should takes one arguments, but handler takes %d", handler.NumIn())
	}
	contextType := reflect.TypeOf((*context.Context)(nil)).Elem()
	argumentType := handler.In(0)
	if !argumentType.Implements(contextType) {
		return fmt.Errorf("handler takes a single arguments, but it is not Context. got %s",
			argumentType.Kind())
	}
	return nil
}

func validateLifeCycleReturns(handler reflect.Type) error {
	if handler.NumOut() != 0 {
		return fmt.Errorf("handler should not have a return value")
	}
	return nil
}

// NewLifeCycleHandler creates a http fc handler from the given handler function.
// The handler function parameter must satisfy the rules documented by RegisterInitializeHandler.
// If handlerFunc is not a valid handler, the returned Handler simply reports the validation error.
func NewLifeCycleHandler(handlerFunc interface{}) LifeCycleHandler {
	if handlerFunc == nil {
		return errorLifeCycleHandler(fmt.Errorf("handler is nil"))
	}
	handler := reflect.ValueOf(handlerFunc)
	handlerType := reflect.TypeOf(handlerFunc)
	if handlerType.Kind() != reflect.Func {
		return errorLifeCycleHandler(fmt.Errorf("handler kind %s is not %s", handlerType.Kind(), reflect.Func))
	}

	err := validateLifeCycleArguments(handlerType)
	if err != nil {
		return errorLifeCycleHandler(err)
	}

	if err := validateLifeCycleReturns(handlerType); err != nil {
		return errorLifeCycleHandler(err)
	}

	return fcLifeCycleHandler(func(ctx context.Context) {
		handler.Call(
			[]reflect.Value{
				reflect.ValueOf(ctx),
			})
	})
}
