package convey

import "github.com/crackcomm/action-test/testing"

type PackageResult struct {
	PackageName string
	Elapsed     float64
	Coverage    float64
	Outcome     string
	BuildOutput string
	TestResults []TestResult
}

func resultPackage(result *testing.Result) *PackageResult {
	pkg := new(PackageResult)
	pkg.PackageName = result.Test.Name
	pkg.Elapsed = result.Duration.Seconds()
	if result.Passed && result.Error == nil {
		pkg.Outcome = "passed"
	} else {
		if result.Error != nil {
			pkg.BuildOutput = result.Error.Error()
		}
		pkg.Outcome = "failed"
	}
	pkg.TestResults = testResults(result)
	return pkg
}

func resultsPackages(results *testing.Results) []*PackageResult {
	packages := []*PackageResult{}
	for _, res := range results.List {
		packages = append(packages, resultPackage(res))
	}
	return flattenPackages(packages)
}

func joinPackages(packages []*PackageResult, pkg *PackageResult) []*PackageResult {
	for _, npkg := range packages {
		if pkg.PackageName == npkg.PackageName {
			npkg.TestResults = append(npkg.TestResults, pkg.TestResults...)
			return packages
		}
	}
	return append(packages, pkg)
}

func flattenPackages(packages []*PackageResult) []*PackageResult {
	result := []*PackageResult{}
	for _, pkg := range packages {
		result = joinPackages(result, pkg)
	}
	return result
}
