package main

import (
	"flag"
	"net/http"
	"strings"

	"github.com/crackcomm/convey-actions/api"
	"github.com/crackcomm/convey-actions/executor"
	"github.com/crackcomm/go-actions/core"
	"github.com/golang/glog"

	_ "github.com/crackcomm/go-actions/source/file"
	_ "github.com/crackcomm/go-actions/source/http"
	_ "github.com/crackcomm/go-core"
)

var (
	listenaddr = "127.0.0.1:8080"
	dashboard  = "convey-dashboard"
	sources    = "actions"
	tests      = "tests"
	watch      = true
)

func init() {
	flag.StringVar(&listenaddr, "listen", listenaddr, "HTTP Listening address")
	flag.StringVar(&dashboard, "dashboard", dashboard, "Dashboard directory")
	flag.StringVar(&sources, "sources", sources, "Actions sources (comma separated)")
	flag.StringVar(&tests, "tests", tests, "Tests directory")
	flag.BoolVar(&watch, "watch", watch, "Watch for changes in tests")
}

func main() {
	defer glog.Flush()
	flag.Set("logtostderr", "true")
	flag.Set("v", "2")
	flag.Parse()

	// Add actions sources
	for _, source := range strings.Split(sources, ",") {
		core.Source(source)
	}

	// Executor
	ex := executor.New(tests)

	// API
	http.Handle("/watch", api.WatchHandler{Executor: ex})
	http.Handle("/execute", api.ExecuteHandler{Executor: ex})
	http.Handle("/status/poll", api.PollHandler{Executor: ex})
	http.Handle("/status", api.StatusHandler{Executor: ex})
	http.Handle("/latest", api.LatestHandler{Executor: ex})

	// Dashboard
	http.Handle("/", http.FileServer(http.Dir(dashboard)))

	if watch {
		go ex.Watch()
	}

	// Start listening
	glog.Infof("Listening on address %s\n", listenaddr)
	glog.Fatal(http.ListenAndServe(listenaddr, nil))
}
