package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type silent struct{}

func (s *silent) Trace(a ...interface{}) {}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func Off() Tracer {
	return &silent{}
}
