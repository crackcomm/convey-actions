package convey

import "time"
import "github.com/crackcomm/action-test/testing"

type CompleteOutput struct {
	Packages []*PackageResult
	Revision string
	Paused   bool
}

func NewResults(results *testing.Results) *CompleteOutput {
	return &CompleteOutput{
		Packages: resultsPackages(results),
		Revision: time.Now().String(),
	}
}
