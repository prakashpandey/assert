package assert

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	cases := []struct {
		name string
		pass bool
		val1 interface{}
		val2 interface{}
	}{
		{
			name: "args_equal",
			pass: true,
			val1: true,
			val2: true,
		},
		{
			name: "args_not_equal",
			pass: false,
			val1: true,
			val2: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var eq = map[bool]string{
				true:  "equal",
				false: "not equal",
			}
			mockT := newMockTester()
			if actual := Equal(c.val1, c.val2, fmt.Sprintf("Expected args to be %s", eq[c.pass]), mockT); actual != c.pass && !mockT.Failed() {
				t.Errorf("Expected test to fail Arg1: %v Arg2: %v\n", c.val1, c.val2)
			}
		})
	}
}

func TestNotEqual(t *testing.T) {
	cases := []struct {
		name string
		pass bool
		val1 interface{}
		val2 interface{}
	}{
		{
			name: "args_not_equal",
			pass: true,
			val1: true,
			val2: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var eq = map[bool]string{
				true:  "not equal",
				false: "equal",
			}
			mockT := newMockTester()
			if actual := NotEqual(c.val1, c.val2, fmt.Sprintf("Expected args to be %s", eq[c.pass]), mockT); actual != c.pass {
				t.Errorf("Expected args to be %s but found: %s", eq[c.pass], eq[!c.pass])
			}
		})
	}
}
