package debug

import (
	"fmt"
	"path/filepath"
)

func ExampleErrorf() {
	err := Errorf("This is a message")
	fmt.Println(
		filepath.Base(err.File()),
		err.Line(),
	)
	// Output:
	// example_test.go 9
}
