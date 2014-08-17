package executor

import "sync"
import "path/filepath"
import "github.com/go-fsnotify/fsnotify"
import "github.com/crackcomm/go-actions/core"
import "github.com/crackcomm/go-actions/action"
import "github.com/crackcomm/action-test/utils"
import "github.com/crackcomm/action-test/testing"
import "github.com/crackcomm/convey-actions/executor/convey"

// NewStatus - Status of a new executor.
var NewStatus = "executing"

// IdleStatus - Status of idle executor.
var IdleStatus = "idle"

// Executor - Takes care of executing tests.
type Executor struct {
	Changed  bool
	Dirname  string
	Status   string
	Tests    testing.Tests
	Actions  *action.Actions
	Watcher  *fsnotify.Watcher
	Latest   *convey.CompleteOutput
	watchers []chan bool
	lock     *sync.RWMutex
}

// New - Creates a new Executor.
func New(dirname string) *Executor {
	dirname, _ = filepath.Abs(dirname)
	return &Executor{
		Changed: true,
		Status:  NewStatus,
		Dirname: dirname,
		lock:    new(sync.RWMutex),
	}
}

// Run - Runs tests and returns convey output straight for API.
func (e *Executor) Run() {
	// Read tests from directory
	tests, _ := utils.ReadTests(e.Dirname)

	e.lock.Lock()
	// Set core actions (cleans up cached actions)
	core.Default.Registry.Cache.Flush()
	// Run tests
	run := tests.Run()
	e.lock.Unlock()

	// Convert to convey output
	e.Latest = convey.NewResults(run)
	e.Status = IdleStatus

	// Run event for API to lookup for changes
	e.RunEvent()
}

// Watch - Creates a loop waiting for events in tests directory and sources.
func (e *Executor) Watch(directories ...string) (err error) {
	e.Watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer e.Watcher.Close()

	// Add all directories including tests to watchlist
	directories = append(directories, e.Dirname)
	for _, dirname := range directories {
		err = e.Watcher.Add(dirname)
		if err != nil {
			return
		}
	}

	// Run watching loop
	for {
		<-e.Watcher.Events
		e.Run()
	}

	return
}

// Events - Returns channel that will return only one event, when something will change, then its closed.
func (e *Executor) Events() chan bool {
	ch := make(chan bool)
	e.watchers = append(e.watchers, ch)
	return ch
}

// RunEvent - Runs event on all watchers watching through <-e.Events().
func (e *Executor) RunEvent() {
	e.Changed = true
	for _, ch := range e.watchers {
		ch <- true
		close(ch)
	}
	e.watchers = []chan bool{}
}
