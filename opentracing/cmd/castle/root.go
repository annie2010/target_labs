// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gopherland/target_labs/opentracing/internal"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	ext "github.com/opentracing/opentracing-go/ext"
	"github.com/spf13/cobra"
)

const (
	app       = "Castle"
	port      = ":5501"
	nightKing = "nightking"
)

var (
	jaeger  string
	rootCmd = &cobra.Command{
		Use:   strings.ToLower(app),
		Short: "G.O.T " + app,
		Long:  "G.O.T " + app,
		Run:   web,
	}
)

// Execute runs the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&jaeger, "jaeger", "j", "localhost:6831", "Specify Jaeger address")
}

func web(cmd *cobra.Command, args []string) {
	closer, err := internal.InitJaeger(app, jaeger)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = closer.Close()
	}()

	r := mux.NewRouter()
	m := handlers.LoggingHandler(os.Stdout, r)
	r.Handle("/api/v1/melt", http.HandlerFunc(meltHandler)).Methods("POST")

	log.Printf("G.O.T %s is listening on port %s...\n", app, port)
	log.Panic(http.ListenAndServe(port, m))
}

func meltHandler(w http.ResponseWriter, r *http.Request) {
	span, err := newSpanFromReq(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	defer span.Finish()

	// creates a span context so we can pass it around.
	ctx := opentracing.ContextWithSpan(r.Context(), span)

	var q internal.Quest
	if q, err = readQuest(ctx, r.Body); err != nil {
		internal.WriteErrOut(ctx, w, err)
		return
	}
	defer func() {
		if r.Body != nil {
			_ = r.Body.Close()
		}
	}()

	span.SetTag("knight", q.Knight)

	log.Printf("Got melt request from %s", q.Knight)
	if !isNightKing(q.Knight) {
		internal.WriteErrOut(ctx, w, fmt.Errorf("💣 Only the `NightKing` can melt this castle!"))
		return
	}

	if err := writeResponse(ctx, w); err == nil {
		span.LogKV("message", fmt.Sprintf("castle successfully melted"))
	}
}

func isNightKing(knight string) bool {
	return strings.ToLower(knight) == nightKing
}

func newSpanFromReq(r *http.Request) (opentracing.Span, error) {
	var span opentracing.Span

	// Extract SpanContext from Headers.
	spanCtx, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		return span, err
	}

	// Starts a child span from the passed in spanContext.
	span, err = opentracing.StartSpan(r.URL.String(), ext.RPCServerOption(spanCtx)), nil
	if err != nil {
		return span, err
	}
	span.SetTag("component", app)
	span.SetTag("http.method", r.Method)
	span.SetTag("http.url", r.URL)

	return span, nil
}

func readQuest(ctx context.Context, body io.ReadCloser) (internal.Quest, error) {
	// Starts a new child span from the given context.
	span := opentracing.StartSpan(
		internal.FuncName(0),
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
	)
	defer span.Finish()

	var q internal.Quest
	if err := json.NewDecoder(body).Decode(&q); err != nil {
		return q, err
	}
	span.SetTag("action", "castle.readQuest")
	span.LogKV("message", fmt.Sprintf("%s requested a melt", q.Knight))

	return q, nil
}

func writeResponse(ctx context.Context, w http.ResponseWriter) error {
	// Starts a new span from the given context.
	span := opentracing.StartSpan(
		internal.FuncName(0),
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
	)
	defer span.Finish()

	resp := internal.Response{Status: "🙀💀 Castle Melted!!"}
	raw, err := json.Marshal(resp)
	if err != nil {
		spanCtx := opentracing.ContextWithSpan(ctx, span)
		internal.WriteErrOut(spanCtx, w, err)
		return err
	}
	span.SetTag("action", "castle.writeResponse")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_, err = w.Write(raw)
	return err
}
