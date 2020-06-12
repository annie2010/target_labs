// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gopherland/target_labs/opentracing/internal"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
)

const (
	port = ":5500"
	app  = "Knight"
)

var (
	jaeger  string
	castle  string
	rootCmd = &cobra.Command{
		Use:   "north",
		Short: "G.O.T " + app + " Service",
		Long:  "G.O.T " + app + " Service",
		Run:   listen,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&jaeger, "jaeger", "j", "localhost:6831", "Specify Jaeger address")
	rootCmd.Flags().StringVarP(&castle, "castle", "c", "localhost:5501", "Specify a castle service address")
}

// Execute a command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func listen(cmd *cobra.Command, args []string) {
	r := mux.NewRouter()
	m := handlers.LoggingHandler(os.Stdout, r)

	closer, err := internal.InitJaeger(app, jaeger)
	defer func() {
		_ = closer.Close()
	}()
	if err != nil {
		panic(err)
	}

	r.Handle("/api/v1/melt", http.HandlerFunc(meltHandler)).Methods("POST")
	log.Printf("G.O.T %s listening on port %s...\n", app, port)
	log.Panic(http.ListenAndServe(port, m))
}

func meltHandler(w http.ResponseWriter, r *http.Request) {
	// Create a top level span from the incoming request.
	span := spanFromReq(r)
	defer span.Finish()
	// Hydrate a context from this span, so we can pass it around.
	spanCtx := opentracing.ContextWithSpan(r.Context(), span)

	var resp internal.Response
	if err := internal.Call(spanCtx, "POST", urlFor("api/v1/melt"), r.Body, &resp); err != nil {
		internal.WriteErrOut(spanCtx, w, err)
		return
	}

	fmt.Println("YO!", resp)
	// writes out the quest response.
	writeResponse(spanCtx, w, resp)
}

func urlFor(path string) string {
	return "http://" + castle + "/" + path
}

// spanFromReq builds a root span from the incoming request.
func spanFromReq(r *http.Request) opentracing.Span {
	span := opentracing.StartSpan(r.URL.String()).
		SetTag("component", app).
		SetTag("http.method", r.Method).
		SetTag("http.url", r.URL)
	span.LogKV("message", internal.FuncName(1))

	return span
}

func writeResponse(ctx context.Context, w http.ResponseWriter, r internal.Response) {
	span := opentracing.StartSpan(
		internal.FuncName(0),
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()),
	)
	defer span.Finish()

	sCtx := opentracing.ContextWithSpan(ctx, span)
	var raw bytes.Buffer
	if err := json.NewEncoder(&raw).Encode(r); err != nil {
		internal.WriteErrOut(sCtx, w, err)
	}
	span.SetTag("http.status_code", http.StatusOK)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(raw.Bytes()); err != nil {
		internal.WriteErrOut(sCtx, w, err)
	}
}
