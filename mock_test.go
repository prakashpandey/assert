package assert

type mockTester struct {
	pass bool
}

func newMockTester() *mockTester {
	return &mockTester{pass: true}
}

func (t *mockTester) Error(args ...interface{}) {
	t.pass = false
}

func (t *mockTester) Errorf(format string, args ...interface{}) {
	t.pass = false
}

func (t *mockTester) Fail() {
	t.pass = false
}

func (t *mockTester) FailNow() {
	t.pass = false
}

func (t *mockTester) Failed() bool {
	// #ReVisit #Bug
	return !t.pass
}

func (t *mockTester) Fatal(args ...interface{}) {
	t.pass = false
}

func (t *mockTester) Fatalf(format string, args ...interface{}) {
	t.pass = false
}

func (t *mockTester) Helper()                                  {}
func (t *mockTester) Log(args ...interface{})                  {}
func (t *mockTester) Logf(format string, args ...interface{})  {}
func (t *mockTester) Name() string                             { return "" }
func (t *mockTester) Parallel()                                {}
func (t *mockTester) Run(name string, f func(t Tester)) bool   { return false }
func (t *mockTester) Skip(args ...interface{})                 {}
func (t *mockTester) SkipNow()                                 {}
func (t *mockTester) Skipf(format string, args ...interface{}) {}
func (t *mockTester) Skipped() bool                            { return false }
