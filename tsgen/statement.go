package tsgen

import (
	"bytes"
	"io"
)

type Statement []Code

func newStatement() *Statement {
	return &Statement{}
}

// Clone makes a copy of the Statement, so further tokens can be appended
// without affecting the original.
func (s *Statement) Clone() *Statement {
	return &Statement{s}
}

func (s *Statement) previous(c Code) Code {
	index := -1
	for i, item := range *s {
		if item == c {
			index = i
			break
		}
	}
	if index > 0 {
		return (*s)[index-1]
	}
	return nil
}

func (s *Statement) isNull(f *File) bool {
	if s == nil {
		return true
	}
	for _, c := range *s {
		if !c.isNull(f) {
			return false
		}
	}
	return true
}

func (s *Statement) render(f *File, w io.Writer, _ *Statement) error {
	first := true
	for _, code := range *s {
		if code == nil || code.isNull(f) {
			// Null() token produces no output but also
			// no separator. Empty() token products no
			// output but adds a separator.
			continue
		}
		if !first {
			if _, err := w.Write([]byte(" ")); err != nil {
				return err
			}
		}
		if err := code.render(f, w, s); err != nil {
			return err
		}
		first = false
	}
	return nil
}

// Render renders the Statement to the provided writer.
func (s *Statement) Render(writer io.Writer) error {
	return s.RenderWithFile(writer, NewFile())
}

// GoString renders the Statement for testing. Any error will cause a panic.
func (s *Statement) TypescriptString() string {
	buf := bytes.Buffer{}
	if err := s.Render(&buf); err != nil {
		panic(err)
	}
	return buf.String()
}

// RenderWithFile renders the Statement to the provided writer, using imports from the provided file.
func (s *Statement) RenderWithFile(writer io.Writer, file *File) error {
	buf := &bytes.Buffer{}
	if err := s.render(file, buf, nil); err != nil {
		return err
	}
	if _, err := writer.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
