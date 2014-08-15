package executor

import "path/filepath"
import "github.com/crackcomm/action-test/utils"
import "github.com/crackcomm/action-test/testing"
import "github.com/crackcomm/convey-actions/executor/convey"

// NewStatus - Status of a new executor.
var NewStatus = "new"

// Executor - Takes care of executing tests.
type Executor struct {
	Dirname string
	Status  string
	Tests   testing.Tests
}

// New - Creates a new Executor.
func New(dirname string) *Executor {
	dirname, _ = filepath.Abs(dirname)
	tests, _ := utils.ReadTests(dirname)
	return &Executor{
		Status:  NewStatus,
		Dirname: dirname,
		Tests:   tests,
	}
}

func (e *Executor) Run() *convey.CompleteOutput {
	return convey.NewResults(e.Tests.Run())
}
