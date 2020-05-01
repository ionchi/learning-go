## Testing in GO

#### Test types

* Test: Unit, Integration, End to End
* Benchmark: Performance
* Example: Documentation

#### Naming conventions

1. Add _test to filenames [main_test.go] - in production builds these files are excluded
2. Prefix tests with "Test" or with "Benchmark" for benchmark tests
3. Accept one parameter - *testing.T or *testing.B for benchmark tests
4. Package naming
    * Same package for whitebox tests
    * Add _test suffix to package for blackbox tests (preferred)
    
#### Reporting Test Failures

* Immediate failure: exit test immediately - for catastrophic failure condition [ex. db connection in an end-to-end test]
    * t.FailNow()
    * t.Fatal(args ...interface{})
    * t.Fatalf(format string, args ...interface{})
* Non-immediate failure: indicate a failure occurred but test function will continue executing
    * t.Fail()
    * t.Error(args ...interface{})
    * t.Errorf(format string, args ...interface{})

#### Test commands

* `go test` - Run all tests in current directory
* `go test {pkg1} {pkg2} ...` - Test specified packages
* `go test ./..` - Run tests in current package descendants
* `go test -v` - Generate verbose output
* `go test -run {regexp}` - Run only tests matching {regexp}
* `go test -cover` - Getting tests coverage
* `go test -coverprofile {output-file}` - generate detailed coverage report
    * then with `go tool cover -func {output-file}` we can read the report
    * or with `go tool cover -html {output-file}` we can have a html output

#### Testing-related Packages

Standard library:
* testing: standard
* testing/quick: design to simplify black box testing
* testing/iotest: provide simple ways to create Readers and Writers that allow to test common 
scenarios when working with input and output
* net/http/httptest: for a network aware applications

Community Projects:
* github.com/stretchrcom/testify
* github.com/onsi/ginkgo
* goconvey.co
* github.com/DATA-DOG/go-sqlmock
