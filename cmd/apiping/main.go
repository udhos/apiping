// Package main implements the apiping tool.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/udhos/otelconfig/oteltrace"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/yaml.v3"
)

const version = "1.3.3"

type application struct {
	me            string
	conf          config
	targets       []string
	server        *http.Server
	serverMetrics *http.Server
	serverHealth  *http.Server
	tracer        trace.Tracer
	met           *metrics
}

func longVersion(me string) string {
	return fmt.Sprintf("%s runtime=%s GOOS=%s GOARCH=%s GOMAXPROCS=%d",
		me, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(0))
}

func main() {

	//
	// parse cmd line
	//

	var showVersion bool
	flag.BoolVar(&showVersion, "version", showVersion, "show version")
	flag.Parse()

	//
	// show version
	//

	me := filepath.Base(os.Args[0])

	{
		v := longVersion(me + " version=" + version)
		if showVersion {
			fmt.Println(v)
			return
		}
		log.Print(v)
	}

	app := &application{
		me:   me,
		conf: getConfig(),
	}

	errTargets := yaml.Unmarshal([]byte(app.conf.targets), &app.targets)
	if errTargets != nil {
		log.Fatalf("error parsing targets: %s: %v", app.conf.targets, errTargets)
	}

	//
	// initialize tracing
	//

	{
		options := oteltrace.TraceOptions{
			DefaultService:     me,
			NoopTracerProvider: false,
			Debug:              true,
		}

		tracer, cancel, errTracer := oteltrace.TraceStart(options)

		if errTracer != nil {
			log.Fatalf("tracer: %v", errTracer)
		}

		defer cancel()

		app.tracer = tracer
	}

	/*
		{
			tp, errTracer := tracerProvider(app.me, app.conf.exporter)
			if errTracer != nil {
				log.Fatalf("tracer provider: %v", errTracer)
			}

			// Register our TracerProvider as the global so any imported
			// instrumentation in the future will default to using it.
			otel.SetTracerProvider(tp)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Cleanly shutdown and flush telemetry when the application exits.
			defer func(ctx context.Context) {
				// Do not make the application hang when it is shutdown.
				ctx, cancel = context.WithTimeout(ctx, time.Second*5)
				defer cancel()
				if err := tp.Shutdown(ctx); err != nil {
					log.Fatalf("trace shutdown: %v", err)
				}
			}(ctx)

			tracePropagation()

			app.tracer = tp.Tracer(fmt.Sprintf("%s-main", app.me))
		}
	*/

	//
	// initialize http
	//

	{
		mux := http.NewServeMux()
		app.server = &http.Server{
			Addr:    app.conf.addr,
			Handler: mux,
		}

		register(mux, app.server.Addr, "handlerRoot", "/", func(w http.ResponseWriter, r *http.Request) { handlerRoot(app, w, r) })
		register(mux, app.server.Addr, "handlerRoute", app.conf.route, func(w http.ResponseWriter, r *http.Request) { handlerRoute(app, w, r) })
	}

	//
	// start http server
	//

	go func() {
		log.Printf("application server: listening on %s", app.conf.addr)
		err := app.server.ListenAndServe()
		log.Fatalf("application server: exited: %v", err)
	}()

	//
	// start metrics server
	//

	{
		app.met = newMetrics(app.conf.metricsNamespace, app.conf.metricsLatencyBucketsServer, app.conf.metricsLatencyBucketsClient)

		mux := http.NewServeMux()
		app.serverMetrics = &http.Server{
			Addr:    app.conf.metricsAddr,
			Handler: mux,
		}

		mux.Handle(app.conf.metricsPath, promhttp.Handler())

		go func() {
			log.Printf("metrics server: listening on %s %s", app.conf.metricsAddr, app.conf.metricsPath)
			err := app.serverMetrics.ListenAndServe()
			log.Fatalf("metrics server: exited: %v", err)
		}()
	}

	//
	// start health server
	//

	{
		mux := http.NewServeMux()
		app.serverHealth = &http.Server{
			Addr:    app.conf.healthAddr,
			Handler: mux,
		}

		mux.HandleFunc(app.conf.healthPath, func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "health ok", 200)
		})

		go func() {
			log.Printf("health server: listening on %s %s", app.conf.healthAddr, app.conf.healthPath)
			err := app.serverHealth.ListenAndServe()
			log.Fatalf("health server: exited: %v", err)
		}()
	}

	//
	// start pinger
	//

	go pinger(app)

	<-make(chan struct{}) // wait forever
}
