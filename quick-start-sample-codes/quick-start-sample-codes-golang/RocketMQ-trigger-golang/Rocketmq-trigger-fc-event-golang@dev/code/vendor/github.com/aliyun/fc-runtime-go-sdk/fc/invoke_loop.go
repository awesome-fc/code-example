// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved
// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package fc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc/messages"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

const (
	msPerS  = int64(time.Second / time.Millisecond)
	nsPerMS = int64(time.Millisecond / time.Nanosecond)
)

type handlerWrapper struct {
	handler  interface{}
	funcType functionType
}

// startRuntimeAPILoop will return an error if handling a particular invoke resulted in a non-recoverable error
// func startRuntimeAPILoop(ctx context.Context, api string, handler interface{}, funcType functionType) error {
func startRuntimeAPILoop(ctx context.Context, api string, baseHandler handlerWrapper, lifeCycleHandlers []handlerWrapper) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("%v", r)
		}
	}()
	client := newRuntimeAPIClient(api)
	function := NewFunction(baseHandler.handler, baseHandler.funcType).withContext(ctx)
	function.RegistryLifeCycleHandler(lifeCycleHandlers)
	for {
		req, err := client.next()
		if err != nil {
			logPrintf("failed to get invoke request due to %v", err)
			continue
		}
		go func(req *invoke, f *Function) {
			err = handleInvoke(req, function)
			if err != nil {
				logPrintf("failed to invoke function due to %v", err)
			}
		}(req, function)
	}
}

// handleInvoke returns an error if the function panics, or some other non-recoverable error occurred
func handleInvoke(invokeInstance *invoke, function *Function) error {
	functionRequest, err := convertInvokeRequest(invokeInstance)
	if err != nil {
		return fmt.Errorf("unexpected error occurred when parsing the invoke: %v", err)
	}

	functionResponse := &messages.InvokeResponse{}
	ivkErr := function.Invoke(functionRequest, functionResponse, convertInvokeFunctionType(invokeInstance))
	if functionResponse.Error != nil {
		payload := safeMarshal(functionResponse.Error)
		if err := invokeInstance.failure(payload, contentTypeJSON); err != nil {
			return fmt.Errorf("unexpected error occurred when sending the function error to the API: %v", err)
		}
		if functionResponse.Error.ShouldExit {
			return fmt.Errorf("calling the handler function resulted in a panic")
		}
		return ivkErr
	}
	if ivkErr != nil {
		return ivkErr
	}

	if err := invokeInstance.success(functionResponse.Payload, contentTypeJSON, functionResponse.HttpParam); err != nil {
		return fmt.Errorf("unexpected error occurred when sending the function functionResponse to the API: %v", err)
	}
	return nil
}

func convertInvokeFunctionType(invokeInstance *invoke) functionType {
	funcType, err := strconv.ParseInt(invokeInstance.headers.Get(headerFunctionType), 10, 64)
	if err != nil {
		return handleFunction
	}
	switch funcType {
	case int64(initializerFunction):
		return initializerFunction
	case int64(preFreezeFunction):
		return preFreezeFunction
	case int64(preStopFunction):
		return preStopFunction
	default:
		return handleFunction
	}

}

// convertInvokeRequest converts an invoke from the Runtime API, and unpacks it to be compatible with the shape of a `lambda.Function` InvokeRequest.
func convertInvokeRequest(invokeInstance *invoke) (*messages.InvokeRequest, error) {
	deadlineEpochMS, err := strconv.ParseInt(invokeInstance.headers.Get(headerDeadlineMS), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse contents of header: %s", headerDeadlineMS)
	}
	deadlineS := deadlineEpochMS / msPerS
	deadlineNS := (deadlineEpochMS % msPerS) * nsPerMS

	functionTimeoutSec, err := strconv.Atoi(invokeInstance.headers.Get(headerFunctionTimeout))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contents of header: %s", headerFunctionTimeout)
	}

	retryCount := 0
	if retryCountStr := invokeInstance.headers.Get(headerRetryCount); retryCountStr != "" {
		retryCount, err = strconv.Atoi(retryCountStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse contents of header: %s", headerFunctionTimeout)
		}
	}

	spanBaggages := make(map[string]string)
	if base64SpanBaggages := invokeInstance.headers.Get(headerOpenTracingSpanBaggages); base64SpanBaggages != "" {
		spanBaggagesByte, err := base64.StdEncoding.DecodeString(base64SpanBaggages)
		if err != nil {
			return nil, fmt.Errorf("failed to parse contents of header %s: %s", headerOpenTracingSpanContext, base64SpanBaggages)
		}
		if err := json.Unmarshal(spanBaggagesByte, &spanBaggages); err != nil {
			return nil, fmt.Errorf("failed to parse contents of header %s: %s", headerOpenTracingSpanContext, base64SpanBaggages)
		}
	}

	res := &messages.InvokeRequest{
		RequestId: invokeInstance.id,
		Deadline: messages.InvokeRequest_Timestamp{
			Seconds: deadlineS,
			Nanos:   deadlineNS,
		},
		Payload: invokeInstance.payload,
		Context: fccontext.FcContext{
			RequestID: invokeInstance.id,
			Credentials: fccontext.Credentials{
				AccessKeyId:     invokeInstance.headers.Get(headerAccessKeyId),
				AccessKeySecret: invokeInstance.headers.Get(headerAccessKeySecret),
				SecurityToken:   invokeInstance.headers.Get(headerSecurityToken),
			},
			Function: fccontext.Function{
				Name:    invokeInstance.headers.Get(headerFunctionName),
				Handler: invokeInstance.headers.Get(headerFunctionHandler),
				Memory:  invokeInstance.headers.Get(headerFunctionMemory),
				Timeout: functionTimeoutSec,
			},
			Service: fccontext.Service{
				Name:       invokeInstance.headers.Get(headerServiceName),
				LogProject: invokeInstance.headers.Get(headerServiceLogproject),
				LogStore:   invokeInstance.headers.Get(headerServiceLogstore),
				Qualifier:  invokeInstance.headers.Get(headerQualifier),
				VersionId:  invokeInstance.headers.Get(headerVersionId),
			},
			Tracing: fccontext.Tracing{
				OpenTracingSpanContext:  invokeInstance.headers.Get(headerOpenTracingSpanContext),
				OpenTracingSpanBaggages: spanBaggages,
				JaegerEndpoint:          invokeInstance.headers.Get(headerJaegerEndpoint),
			},
			Region:     invokeInstance.headers.Get(headerRegion),
			AccountId:  invokeInstance.headers.Get(headerAccountId),
			RetryCount: retryCount,
		},
	}

	if httpParams := invokeInstance.headers.Get(headerHttpParams); httpParams != "" {
		res.HttpParams = &httpParams
	}

	return res, nil
}

func safeMarshal(v interface{}) []byte {
	payload, err := json.Marshal(v)
	if err != nil {
		v := &messages.InvokeResponse_Error{
			Type:    "Runtime.SerializationError",
			Message: err.Error(),
		}
		payload, err := json.Marshal(v)
		if err != nil {
			panic(err) // never reach
		}
		return payload
	}
	return payload
}
