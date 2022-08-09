// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.
//
// Helpers for accessing context information from an Invoke request. Context information
// is stored in a https://golang.org/pkg/context/#Context. The functions FromContext
// are used to retrieving an instance of FcContext.

package fccontext

import "context"

var FunctionVersion string

// Credentials ...
type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

// Function ...
type Function struct {
	Name    string
	Handler string
	Memory  string
	Timeout int
}

// Service ...
type Service struct {
	Name       string
	LogProject string
	LogStore   string
	Qualifier  string
	VersionId  string
}

// Tracing ...
type Tracing struct {
	OpenTracingSpanContext  string
	OpenTracingSpanBaggages map[string]string
	JaegerEndpoint          string
}

// FcContext is the set of metadata that is passed for every Invoke.
type FcContext struct {
	RequestID   string
	Credentials Credentials
	Function    Function
	Service     Service
	Region      string
	AccountId   string
	RetryCount  int
	Tracing     Tracing

	logger *FcLogger
}

func (f *FcContext) GetLogger() *FcLogger {
	return f.logger
}

// An unexported type to be used as the key for types in this package.
// This prevents collisions with keys defined in other packages.
type key struct{}

// The key for a LambdaContext in Contexts.
// Users of this package must use lambdacontext.NewContext and lambdacontext.FromContext
// instead of using this key directly.
var contextKey = &key{}

// NewContext returns a new Context that carries value lc.
func NewContext(parent context.Context, lc *FcContext) context.Context {
	return context.WithValue(parent, contextKey, lc)
}

// FromContext returns the LambdaContext value stored in ctx, if any.
func FromContext(ctx context.Context) (*FcContext, bool) {
	lc, ok := ctx.Value(contextKey).(*FcContext)
	if !ok {
		return nil, false
	}
	lc.logger = NewFcLogger(lc.RequestID)
	return lc, ok
}
