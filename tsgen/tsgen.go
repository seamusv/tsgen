package tsgen

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type Code interface {
	render(f *File, w io.Writer, s *Statement) error
	isNull(f *File) bool
}

func (f *File) Save(filename string) error {
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

func (f *File) Render(w io.Writer) error {
	body := &bytes.Buffer{}
	if err := f.render(f, body, nil); err != nil {
		return err
	}

	source := &bytes.Buffer{}
	if len(f.headers) > 0 {
		for _, header := range f.headers {
			if _, err := fmt.Fprintf(source, "// %s\n", header); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintf(source, "\n"); err != nil {
			return err
		}
	}

	if err := f.renderImports(source); err != nil {
		return err
	}

	if _, err := source.Write(body.Bytes()); err != nil {
		return err
	}

	formatted, err := f.formatter(source.Bytes())
	if err != nil {
		return err
	}

	if _, err := w.Write(formatted); err != nil {
		return err
	}

	return nil
}

func (f *File) renderImports(w io.Writer) error {
	sortedPaths := make([]string, 0)
	{
		for path, _ := range f.imports {
			sortedPaths = append(sortedPaths, path)
		}
		sort.Strings(sortedPaths)
	}

	for _, path := range sortedPaths {
		def := f.imports[path]
		if def.isDefault {
			if def.name == "" {
				if _, err := fmt.Fprintf(w, "import '%s';\n", path); err != nil {
					return err
				}
			} else {
				if _, err := fmt.Fprintf(w, "import %s from '%s';\n", def.name, path); err != nil {
					return err
				}
			}
		} else {
			items := make([]string, 0)
			for _, module := range def.modules {
				alias := ""
				if module.alias != "" {
					alias = fmt.Sprintf(" as %s", module.alias)
				}
				items = append(items, fmt.Sprintf("%s%s", module.name, alias))
			}
			sort.Strings(items)
			if _, err := fmt.Fprintf(w, "import { %s } from '%s';\n", strings.Join(items, ", "), path); err != nil {
				return err
			}
		}
	}

	//if len(f.imports) > 0 {
	//	if _, err := fmt.Fprintf(w, "\n"); err != nil {
	//		return err
	//	}
	//}

	return nil
}
