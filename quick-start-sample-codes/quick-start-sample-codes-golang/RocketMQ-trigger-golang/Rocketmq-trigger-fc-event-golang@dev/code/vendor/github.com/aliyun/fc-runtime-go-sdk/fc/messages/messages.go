// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package messages

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

type PingRequest struct {
}

type PingResponse struct {
}

type InvokeRequest_Timestamp struct {
	Seconds int64
	Nanos   int64
}

type InvokeRequest struct {
	Payload    []byte
	RequestId  string
	Deadline   InvokeRequest_Timestamp
	Context    fccontext.FcContext
	HttpParams *string
}

type InvokeResponse struct {
	Payload []byte
	Error   *InvokeResponse_Error

	// HttpHandler parameter
	HttpParam string
}

type InvokeResponse_Error struct {
	Message    string                             `json:"errorMessage"`
	Type       string                             `json:"errorType"`
	StackTrace []*InvokeResponse_Error_StackFrame `json:"stackTrace,omitempty"`
	ShouldExit bool                               `json:"-"`
}

func (e InvokeResponse_Error) Error() string {
	return fmt.Sprintf("%#v", e)
}

func (e InvokeResponse_Error) ToJson() string {
	res, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		log.Fatalln(err)
	}
	return string(res)
	// return fmt.Sprintf("%#v", e)
}

type InvokeResponse_Error_StackFrame struct {
	Path  string `json:"path"`
	Line  int32  `json:"line"`
	Label string `json:"label"`
}
