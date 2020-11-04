// Package hello_test is external test package for test package.
package hello_test

import (
	"fmt"
	"testing"

	"github.com/devlights/go-external-test-package-example/pkg/hello"
)

func ExampleSay() {
	fmt.Println(hello.Say("golang"))

	// Output:
	// hello golang
}

func TestSay(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{"not empty", "golang", "hello golang"},
		{"empty", "", ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ans := hello.Say(c.in)
			if c.out != ans {
				t.Errorf("[want] %s\t[got] %s", c.out, ans)
			}
		})
	}
}
