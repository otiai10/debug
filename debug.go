package debug

import (
	"fmt"
	"log"
)

// Fatalln ...
func Fatalln(v ...interface{}) {
	v = append([]interface{}{newCaller(DefaultSkip).String()}, v...)
	log.Fatalln(v...)
}

// Errorf ...
func Errorf(format string, v ...interface{}) *Error {
	return &Error{
		Caller: newCaller(DefaultSkip),
		Format: format,
		Values: v,
	}
}

// Error ...
type Error struct {
	Caller *caller
	Format string
	Values []interface{}
}

// Error to satisfy error interface.
func (e *Error) Error() string {
	format := e.Caller.String() + " " + e.Format
	return fmt.Sprintf(format, e.Values...)
}

// File ...
func (e *Error) File() string {
	return e.Caller.File
}

// Line ...
func (e Error) Line() int {
	return e.Caller.Line
}
