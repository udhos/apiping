package main

import (
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
)

type handler struct {
	f http.HandlerFunc
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.f(w, r)
}

func register(mux *http.ServeMux, operation, addr, path string, handlerFunc http.HandlerFunc) {
	h := &handler{f: handlerFunc}
	mux.Handle(path, otelhttp.NewHandler(h, operation))
	log.Printf("registered %s on port %s path %s", operation, addr, path)
}

func handlerRoot(app *application, w http.ResponseWriter, r *http.Request) {
	const me = "handlerRoot"

	_, span := app.tracer.Start(r.Context(), me)
	defer span.End()

	traceID := span.SpanContext().TraceID().String()

	log.Printf("%s: traceID=%s: %s %s %s - 404 not found",
		me, traceID, r.RemoteAddr, r.Method, r.RequestURI)

	span.SetStatus(codes.Error, "404 not found")

	http.Error(w, "not found", 404)
}

func handlerRoute(app *application, w http.ResponseWriter, r *http.Request) {
	const me = "handlerRoute"

	begin := time.Now()

	_, span := app.tracer.Start(r.Context(), me)
	defer span.End()

	traceID := span.SpanContext().TraceID().String()

	elap := time.Since(begin)

	log.Printf("%s: traceID=%s: %s %s %s - 200 ok - elapsed:%v",
		me, traceID, r.RemoteAddr, r.Method, r.RequestURI, elap)

	app.met.recordLatencyServer(r.Method, "200", r.RequestURI, elap)

	http.Error(w, "ok", 200)
}
