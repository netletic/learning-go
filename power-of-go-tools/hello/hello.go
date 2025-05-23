package hello

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func (p Printer) Print() {
	fmt.Fprintln(p.Output, "Hello, world")
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
	}
}

func Main() {
	NewPrinter().Print()
}
