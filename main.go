package main

import "flag"
import "net/http"
import "github.com/golang/glog"
import "github.com/crackcomm/go-actions/core"
import "github.com/crackcomm/convey-actions/api"
import "github.com/crackcomm/convey-actions/executor"
import "github.com/crackcomm/go-actions/source/utils"
import _ "github.com/crackcomm/go-core"

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
	flag.Parse()

	// Add actions sources
	core.AddSources(utils.GetSources(sources)...)

	// Executor
	ex := executor.New(tests)

	// API
	http.Handle("/watch", &api.WatchHandler{ex})
	http.Handle("/execute", &api.ExecuteHandler{ex})
	http.Handle("/status/poll", &api.PollHandler{ex})
	http.Handle("/status", &api.StatusHandler{ex})
	http.Handle("/latest", &api.LatestHandler{ex})

	// Dashboard
	http.Handle("/", http.FileServer(http.Dir(dashboard)))

	if watch {
		go ex.Watch()
	}

	// Start listening
	glog.Infof("Listening on address %s\n", listenaddr)
	glog.Fatal(http.ListenAndServe(listenaddr, nil))
}
