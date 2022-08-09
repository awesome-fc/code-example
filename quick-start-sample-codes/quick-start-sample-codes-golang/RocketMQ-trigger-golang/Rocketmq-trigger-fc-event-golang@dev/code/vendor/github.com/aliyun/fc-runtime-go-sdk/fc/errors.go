// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved
// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package fc

import (
	"reflect"

	"github.com/aliyun/fc-runtime-go-sdk/fc/messages"
)

func getErrorType(err interface{}) string {
	errorType := reflect.TypeOf(err)
	if errorType.Kind() == reflect.Ptr {
		return errorType.Elem().Name()
	}
	return errorType.Name()
}

func fcErrorResponse(invokeError error) *messages.InvokeResponse_Error {
	if ive, ok := invokeError.(messages.InvokeResponse_Error); ok {
		return &ive
	}
	var errorName string
	if errorType := reflect.TypeOf(invokeError); errorType.Kind() == reflect.Ptr {
		errorName = errorType.Elem().Name()
	} else {
		errorName = errorType.Name()
	}
	return &messages.InvokeResponse_Error{
		Message: invokeError.Error(),
		Type:    errorName,
	}
}

func fcPanicResponse(err interface{}) *messages.InvokeResponse_Error {
	if ive, ok := err.(messages.InvokeResponse_Error); ok {
		return &ive
	}
	panicInfo := getPanicInfo(err)
	return &messages.InvokeResponse_Error{
		Message:    panicInfo.Message,
		Type:       getErrorType(err),
		StackTrace: panicInfo.StackTrace,
		ShouldExit: true,
	}
}
