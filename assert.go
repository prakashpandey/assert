package assert

// Tester ...
type Tester interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Parallel()
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

// Equal ...
func Equal(actual, expected interface{}, msg string, t Tester) bool {
	if actual != expected {
		t.Errorf(msg)
		return false
	}
	return true
}

// Equalf ...
func Equalf(actual, expected interface{}, t Tester, format string, args ...interface{}) bool {
	if actual != expected {
		t.Errorf(format, args...)
		return false
	}
	return true
}

// NotEqual ...
func NotEqual(actual, expected interface{}, msg string, t Tester) bool {
	if actual == expected {
		t.Errorf(msg)
		return false
	}
	return true
}

// NotEqualf ...
func NotEqualf(actual, expected interface{}, t Tester, format string, args ...interface{}) bool {
	if actual == expected {
		t.Errorf(format, args...)
		return false
	}
	return true
}
