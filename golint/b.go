package lint

import (
	"errors"
	F "fmt"
	_ "image/gif"
)

var SomeError = errors.New("Capitalised error message")

type unexported int
type Exported int

func (this unexported) Foo() {}

func (oneName Exported) Foo() {
}

func (anotherName Exported) Bar() {
	F.Println("Hi")
}