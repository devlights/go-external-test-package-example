package hello2_test

import (
	"fmt"
	"testing"

	"github.com/devlights/go-external-test-package-example/pkg/hello2"
)

func ExampleSay() {
	fmt.Println(hello2.Say("golang"))

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
			ans := hello2.Say(c.in)
			if c.out != ans {
				t.Errorf("[want] %s\t[got] %s", c.out, ans)
			}
		})
	}
}

func TestMakeMessage(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{"not empty", "golang", "hello golang"},
		{"empty", "", "hello "},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ans := hello2.MakeMessage(c.in)
			if c.out != ans {
				t.Errorf("[want] %s\t[got] %s", c.out, ans)
			}
		})
	}
}