// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
)

type (
	// Response a generic Quest response
	Response struct {
		Status string `json:"status"`
	}

	// Quest represents a knight's quest
	Quest struct {
		Knight string `json:"knight"`
	}
)

// Call a remote http service
func Call(ctx context.Context, method, url string, payload io.ReadCloser, res interface{}) error {
	defer func() {
		if payload != nil {
			_ = payload.Close()
		}
	}()

	// Extract top level span.
	parentSpan := opentracing.SpanFromContext(ctx)

	// Create a new child span for the out request.
	span := opentracing.StartSpan(FuncName(0), opentracing.ChildOf(parentSpan.Context()))
	defer span.Finish()
	span.SetTag("http.method", method)
	span.SetTag("http.url", url)
	spanCtx := opentracing.ContextWithSpan(ctx, span)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return SpanError(spanCtx, err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Ensure we have continuation thru the castle service.
	// Injects the spanContext into the request headers.
	if err = opentracing.GlobalTracer().Inject(
		parentSpan.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	); err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return SpanError(spanCtx, fmt.Errorf("remote call failed %s", err))
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		raw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return SpanError(spanCtx, err)
		}
		return SpanError(spanCtx, fmt.Errorf("call `%s failed (%d) -- %s", url, resp.StatusCode, raw))
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return SpanError(spanCtx, err)
	}

	return nil
}

// SpanError decorates span with the given error
func SpanError(ctx context.Context, err error) error {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("error", true)
	span.LogKV(
		"event", "error",
		"message", err.Error(),
	)

	return err
}

// WriteErrOut formulate err response and decorate span
func WriteErrOut(ctx context.Context, w http.ResponseWriter, err error) {
	_ = SpanError(ctx, err)

	// Extract span from context and annotate it.
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("http.status_code", http.StatusExpectationFailed)

	http.Error(w, err.Error(), http.StatusExpectationFailed)
}
