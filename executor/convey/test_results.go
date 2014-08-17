package convey

import "fmt"
import "github.com/crackcomm/action-test/testing"

type TestResult struct {
	TestName string
	Coverage float64
	Elapsed  float64
	Passed   bool
	Message  string
	Error    string
	Stories  []ScopeResult
}

type ScopeResult struct {
	Title      string
	Assertions []*AssertionResult
}

type AssertionResult struct {
	Expected string
	Actual   string
	Failure  string
	Error    interface{}
}

func variableAssertion(variable *testing.Variable) *AssertionResult {
	assertion := &AssertionResult{}
	assertion.Actual = fmt.Sprintf("%v", variable.Value)
	assertion.Expected = fmt.Sprintf("%v", variable.Expected)
	if !variable.IsExpected() {
		assertion.Failure = fmt.Sprintf("Unexpected context variable %s", variable.Name)
	}
	return assertion
}

func scopeResults(result *testing.Result) []ScopeResult {
	list := []ScopeResult{}
	for name, variable := range result.Variables {
		res := ScopeResult{}
		res.Title = fmt.Sprintf("Context value %s should match", name)
		res.Assertions = []*AssertionResult{variableAssertion(variable)}
		list = append(list, res)
	}
	return list
}

func testResults(result *testing.Result) []TestResult {
	res := TestResult{
		TestName: result.Test.Description,
		Elapsed:  result.Duration.Seconds(),
		Passed:   result.Passed && result.Error == nil,
		Stories:  []ScopeResult{},
	}
	if !result.Passed {
		res.Coverage = -1
		if result.Error != nil {
			res.Error = result.Error.Error()
			res.Message = res.Error
		} else {
			res.Message = fmt.Sprintf("Error at %v", result.Duration)
			res.Stories = scopeResults(result)
		}
	} else {
		res.Message = fmt.Sprintf("Done in %v", result.Duration)
		res.Stories = scopeResults(result)
	}
	return []TestResult{res}
}
