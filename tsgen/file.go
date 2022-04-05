package tsgen

import (
	"bytes"
	"fmt"
)

type Option func(f *File) error

func WithPrettierFormatter() Option {
	return func(f *File) error {
		f.formatter = Prettier
		return nil
	}
}

func NewFile(options ...Option) *File {
	f := &File{
		Group: &Group{
			multi: true,
		},
		imports: map[string]importDef{},
	}

	for _, option := range options {
		if err := option(f); err != nil {
			panic(err)
		}
	}

	if f.formatter == nil {
		f.formatter = func(b []byte) ([]byte, error) {
			return b, nil
		}
	}

	return f
}

func (f *File) HeaderComment(comment string) {
	f.headers = append(f.headers, comment)
}

func (f *File) ImportDefault(path, name string) {
	f.imports[path] = importDef{
		name:      name,
		isDefault: true,
	}
}

func (f *File) Import(path, name string) {
	f.ImportAlias(path, name, "")
}

func (f *File) ImportAlias(path, name, alias string) {
	var modules []importModule
	if m, ok := f.imports[path]; ok {
		modules = m.modules
	}
	modules = append(modules, importModule{
		name:  name,
		alias: alias,
	})
	f.imports[path] = importDef{
		name:    name,
		modules: modules,
	}
}

func (f *File) GoString() string {
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}

	b, err := f.formatter(buf.Bytes())
	if err != nil {
		fmt.Println("Code:")
		fmt.Println(buf.String())
		panic(err)
	}

	return string(b)
}

type File struct {
	*Group
	headers   []string
	imports   map[string]importDef
	formatter func([]byte) ([]byte, error)
}

type importDef struct {
	name      string
	isDefault bool
	modules   []importModule
}

type importModule struct {
	name  string
	alias string
}
