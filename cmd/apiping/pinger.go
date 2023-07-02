package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func pinger(app *application) {
	const me = "pinger"
	for {
		for _, target := range app.targets {
			go pingTarget(target, app.tracer, app.met)
			log.Printf("%s: sleeping for %v", me, app.conf.interval)
		}
		time.Sleep(app.conf.interval)
	}
}

func pingTarget(target string, tracer trace.Tracer, met *metrics) {
	const me = "pingTarget"
	ctx, span := tracer.Start(context.Background(), me)
	defer span.End()

	traceID := span.SpanContext().TraceID()

	const (
		timeout = 5 * time.Second
		method  = "GET"
	)

	log.Printf("%s: URL=%s traceID=%s timeout=%v", me, target, traceID, timeout)

	var status int

	begin := time.Now()

	defer func() {
		elap := time.Since(begin)
		log.Printf("%s: URL=%s traceID=%s elapsed=%v", me, target, traceID, elap)
		met.recordLatencyClient(method, fmt.Sprint(status), target, elap)
	}()

	req, errReq := http.NewRequestWithContext(ctx, method, target, nil)
	if errReq != nil {
		log.Printf("%s: URL=%s traceID=%s request error: %v", me, target, traceID, errReq)
		span.SetStatus(codes.Error, errReq.Error())
		return
	}

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
		Timeout:   timeout,
	}

	resp, errGet := client.Do(req)
	if errGet != nil {
		log.Printf("%s: URL=%s traceID=%s server error: %v", me, target, traceID, errGet)
		span.SetStatus(codes.Error, errGet.Error())
		return
	}

	status = resp.StatusCode // save status for defer

	defer resp.Body.Close()

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		log.Printf("%s: URL=%s traceID=%s body error: %v", me, target, traceID, errBody)
		span.SetStatus(codes.Error, errBody.Error())
		return
	}

	str := string(body)

	if resp.StatusCode != 200 {
		log.Printf("%s: URL=%s traceID=%s bad response status: status=%d %v", me, target, traceID, status, str)
		span.SetStatus(codes.Error, fmt.Sprintf("bad response status: %d", status))
		return
	}

	log.Printf("%s: URL=%s traceID=%s response: status=%d %v", me, target, traceID, status, str)
}
