package hello

import (
	"testing"
)

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
			ans := makeMessage(c.in)
			if c.out != ans {
				t.Errorf("[want] %s\t[got] %s", c.out, ans)
			}
		})
	}
}
