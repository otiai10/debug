package debug

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestNewCaller(t *testing.T) {
	c := newCaller(1)
	Expect(t, c.String()).Match("/all_test.go:10")

	When(t, "The caller is nil", func(t *testing.T) {
		var c *caller
		Expect(t, c.String()).ToBe("")
	})
}

func TestErrorf(t *testing.T) {
	err := Errorf("This is message")
	Expect(t, err.Error()).Match("/all_test.go:20 This is message")
}

func TestFatalln(t *testing.T) {
	Expect(t, func() {
		Fatalln("test test")
	}).Exit(1)
}
