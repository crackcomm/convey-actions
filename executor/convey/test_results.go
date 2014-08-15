package convey

import "fmt"
import "github.com/crackcomm/action-test/testing"

type TestResult struct {
	TestName string
	Elapsed  float64
	Passed   bool
	Skipped  bool
	// File     string
	// Line     int
	// Message  string
	Error    string
	Stories  []ScopeResult
}

type ScopeResult struct {
	Title      string
	Assertions []*AssertionResult
}

type AssertionResult struct {
	Expected   string
	Actual     string
	Failure    string
	Error      interface{}
}

func variableAssertion(variable *testing.Variable) *AssertionResult {
	assertion := &AssertionResult{}
	assertion.Actual = fmt.Sprintf("%v", variable.Value)
	assertion.Expected = fmt.Sprintf("%v", variable.Expected)
	if !variable.IsExpected() {
		assertion.Failure = "TODO FAILURE..."
	}
	return assertion
}

func scopeResults(result *testing.Result) []ScopeResult {
	list := []ScopeResult{}
	for name, variable := range result.Variables {
		res := ScopeResult{}
		res.Title = name
		res.Assertions = []*AssertionResult{variableAssertion(variable)}
	}
	return list
}

func testResults(result *testing.Result) []TestResult {
	res := TestResult{
		TestName: result.Test.Description,
		Elapsed:  result.Duration.Seconds(),
		Passed:   result.Passed && result.Error == nil,
		Skipped:  false,
		Stories:  scopeResults(result),
	}
	if result.Error != nil {
		res.Error = result.Error.Error()
	}
	return []TestResult{res}
}
