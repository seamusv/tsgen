// This file is generated - do not edit.

package tsgen

// Semicolon renders a semicolon
func Semicolon() *Statement {
	return newStatement().Semicolon()
}

// Semicolon renders a semicolon
func (g *Group) Semicolon() *Statement {
	s := Semicolon()
	g.items = append(g.items, s)
	return s
}

// Semicolon renders a semicolon
func (s *Statement) Semicolon() *Statement {
	g := &Group{
		close:     ";",
		items:     []Code{},
		multi:     false,
		name:      "semicolon",
		open:      "",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// Parens renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func Parens(item Code) *Statement {
	return newStatement().Parens(item)
}

// Parens renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func (g *Group) Parens(item Code) *Statement {
	s := Parens(item)
	g.items = append(g.items, s)
	return s
}

// Parens renders a single item in parenthesis. Use for type conversion or to specify evaluation order.
func (s *Statement) Parens(item Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{item},
		multi:     false,
		name:      "parens",
		open:      "(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// List renders a comma separated list. Use for multiple return functions.
func List(items ...Code) *Statement {
	return newStatement().List(items...)
}

// List renders a comma separated list. Use for multiple return functions.
func (g *Group) List(items ...Code) *Statement {
	s := List(items...)
	g.items = append(g.items, s)
	return s
}

// List renders a comma separated list. Use for multiple return functions.
func (s *Statement) List(items ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     items,
		multi:     false,
		name:      "list",
		open:      "",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// ListFunc renders a comma separated list. Use for multiple return functions.
func ListFunc(f func(*Group)) *Statement {
	return newStatement().ListFunc(f)
}

// ListFunc renders a comma separated list. Use for multiple return functions.
func (g *Group) ListFunc(f func(*Group)) *Statement {
	s := ListFunc(f)
	g.items = append(g.items, s)
	return s
}

// ListFunc renders a comma separated list. Use for multiple return functions.
func (s *Statement) ListFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "list",
		open:      "",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Values renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func Values(values ...Code) *Statement {
	return newStatement().Values(values...)
}

// Values renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func (g *Group) Values(values ...Code) *Statement {
	s := Values(values...)
	g.items = append(g.items, s)
	return s
}

// Values renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func (s *Statement) Values(values ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     values,
		multi:     false,
		name:      "values",
		open:      "{",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// ValuesFunc renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func ValuesFunc(f func(*Group)) *Statement {
	return newStatement().ValuesFunc(f)
}

// ValuesFunc renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func (g *Group) ValuesFunc(f func(*Group)) *Statement {
	s := ValuesFunc(f)
	g.items = append(g.items, s)
	return s
}

// ValuesFunc renders a comma separated list enclosed by curly braces. Use for slice or composite literals.
func (s *Statement) ValuesFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		multi:     false,
		name:      "values",
		open:      "{",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Method renders a method call. Use for method calls.
func Method(name Code) *Statement {
	return newStatement().Method(name)
}

// Method renders a method call. Use for method calls.
func (g *Group) Method(name Code) *Statement {
	s := Method(name)
	g.items = append(g.items, s)
	return s
}

// Method renders a method call. Use for method calls.
func (s *Statement) Method(name Code) *Statement {
	g := &Group{
		close:     "",
		items:     []Code{name},
		multi:     false,
		name:      "method",
		open:      "",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// Index renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func Index(items ...Code) *Statement {
	return newStatement().Index(items...)
}

// Index renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func (g *Group) Index(items ...Code) *Statement {
	s := Index(items...)
	g.items = append(g.items, s)
	return s
}

// Index renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func (s *Statement) Index(items ...Code) *Statement {
	g := &Group{
		close:     "]",
		items:     items,
		multi:     false,
		name:      "index",
		open:      "[",
		separator: ":",
	}
	*s = append(*s, g)
	return s
}

// IndexFunc renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func IndexFunc(f func(*Group)) *Statement {
	return newStatement().IndexFunc(f)
}

// IndexFunc renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func (g *Group) IndexFunc(f func(*Group)) *Statement {
	s := IndexFunc(f)
	g.items = append(g.items, s)
	return s
}

// IndexFunc renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.
func (s *Statement) IndexFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "]",
		multi:     false,
		name:      "index",
		open:      "[",
		separator: ":",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Block renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func Block(statements ...Code) *Statement {
	return newStatement().Block(statements...)
}

// Block renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func (g *Group) Block(statements ...Code) *Statement {
	s := Block(statements...)
	g.items = append(g.items, s)
	return s
}

// Block renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func (s *Statement) Block(statements ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     statements,
		multi:     true,
		name:      "block",
		open:      "{",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func BlockFunc(f func(*Group)) *Statement {
	return newStatement().BlockFunc(f)
}

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func (g *Group) BlockFunc(f func(*Group)) *Statement {
	s := BlockFunc(f)
	g.items = append(g.items, s)
	return s
}

// BlockFunc renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.
func (s *Statement) BlockFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		multi:     true,
		name:      "block",
		open:      "{",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Defs renders a statement list enclosed in parenthesis. Use for definition lists.
func Defs(definitions ...Code) *Statement {
	return newStatement().Defs(definitions...)
}

// Defs renders a statement list enclosed in parenthesis. Use for definition lists.
func (g *Group) Defs(definitions ...Code) *Statement {
	s := Defs(definitions...)
	g.items = append(g.items, s)
	return s
}

// Defs renders a statement list enclosed in parenthesis. Use for definition lists.
func (s *Statement) Defs(definitions ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     definitions,
		multi:     true,
		name:      "defs",
		open:      "(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// DefsFunc renders a statement list enclosed in parenthesis. Use for definition lists.
func DefsFunc(f func(*Group)) *Statement {
	return newStatement().DefsFunc(f)
}

// DefsFunc renders a statement list enclosed in parenthesis. Use for definition lists.
func (g *Group) DefsFunc(f func(*Group)) *Statement {
	s := DefsFunc(f)
	g.items = append(g.items, s)
	return s
}

// DefsFunc renders a statement list enclosed in parenthesis. Use for definition lists.
func (s *Statement) DefsFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     true,
		name:      "defs",
		open:      "(",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Call renders a comma separated list enclosed by parenthesis. Use for function calls.
func Call(params ...Code) *Statement {
	return newStatement().Call(params...)
}

// Call renders a comma separated list enclosed by parenthesis. Use for function calls.
func (g *Group) Call(params ...Code) *Statement {
	s := Call(params...)
	g.items = append(g.items, s)
	return s
}

// Call renders a comma separated list enclosed by parenthesis. Use for function calls.
func (s *Statement) Call(params ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     params,
		multi:     false,
		name:      "call",
		open:      "(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// CallFunc renders a comma separated list enclosed by parenthesis. Use for function calls.
func CallFunc(f func(*Group)) *Statement {
	return newStatement().CallFunc(f)
}

// CallFunc renders a comma separated list enclosed by parenthesis. Use for function calls.
func (g *Group) CallFunc(f func(*Group)) *Statement {
	s := CallFunc(f)
	g.items = append(g.items, s)
	return s
}

// CallFunc renders a comma separated list enclosed by parenthesis. Use for function calls.
func (s *Statement) CallFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     false,
		name:      "call",
		open:      "(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Params renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func Params(params ...Code) *Statement {
	return newStatement().Params(params...)
}

// Params renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func (g *Group) Params(params ...Code) *Statement {
	s := Params(params...)
	g.items = append(g.items, s)
	return s
}

// Params renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func (s *Statement) Params(params ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     params,
		multi:     false,
		name:      "params",
		open:      "(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// ParamsFunc renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func ParamsFunc(f func(*Group)) *Statement {
	return newStatement().ParamsFunc(f)
}

// ParamsFunc renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func (g *Group) ParamsFunc(f func(*Group)) *Statement {
	s := ParamsFunc(f)
	g.items = append(g.items, s)
	return s
}

// ParamsFunc renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.
func (s *Statement) ParamsFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     false,
		name:      "params",
		open:      "(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Assert renders a period followed by a single item enclosed by parenthesis. Use for type assertions.
func Assert(typ Code) *Statement {
	return newStatement().Assert(typ)
}

// Assert renders a period followed by a single item enclosed by parenthesis. Use for type assertions.
func (g *Group) Assert(typ Code) *Statement {
	s := Assert(typ)
	g.items = append(g.items, s)
	return s
}

// Assert renders a period followed by a single item enclosed by parenthesis. Use for type assertions.
func (s *Statement) Assert(typ Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{typ},
		multi:     false,
		name:      "assert",
		open:      ".(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// If renders the keyword followed by a semicolon separated list.
func If(conditions ...Code) *Statement {
	return newStatement().If(conditions...)
}

// If renders the keyword followed by a semicolon separated list.
func (g *Group) If(conditions ...Code) *Statement {
	s := If(conditions...)
	g.items = append(g.items, s)
	return s
}

// If renders the keyword followed by a semicolon separated list.
func (s *Statement) If(conditions ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     conditions,
		multi:     false,
		name:      "if",
		open:      "if ",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// IfFunc renders the keyword followed by a semicolon separated list.
func IfFunc(f func(*Group)) *Statement {
	return newStatement().IfFunc(f)
}

// IfFunc renders the keyword followed by a semicolon separated list.
func (g *Group) IfFunc(f func(*Group)) *Statement {
	s := IfFunc(f)
	g.items = append(g.items, s)
	return s
}

// IfFunc renders the keyword followed by a semicolon separated list.
func (s *Statement) IfFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "if",
		open:      "if ",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Return renders the keyword followed by a comma separated list.
func Return(result Code) *Statement {
	return newStatement().Return(result)
}

// Return renders the keyword followed by a comma separated list.
func (g *Group) Return(result Code) *Statement {
	s := Return(result)
	g.items = append(g.items, s)
	return s
}

// Return renders the keyword followed by a comma separated list.
func (s *Statement) Return(result Code) *Statement {
	g := &Group{
		close:     "",
		items:     []Code{result},
		multi:     false,
		name:      "return",
		open:      "return ",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// For renders the keyword followed by a semicolon separated list.
func For(conditions ...Code) *Statement {
	return newStatement().For(conditions...)
}

// For renders the keyword followed by a semicolon separated list.
func (g *Group) For(conditions ...Code) *Statement {
	s := For(conditions...)
	g.items = append(g.items, s)
	return s
}

// For renders the keyword followed by a semicolon separated list.
func (s *Statement) For(conditions ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     conditions,
		multi:     false,
		name:      "for",
		open:      "for ",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// ForFunc renders the keyword followed by a semicolon separated list.
func ForFunc(f func(*Group)) *Statement {
	return newStatement().ForFunc(f)
}

// ForFunc renders the keyword followed by a semicolon separated list.
func (g *Group) ForFunc(f func(*Group)) *Statement {
	s := ForFunc(f)
	g.items = append(g.items, s)
	return s
}

// ForFunc renders the keyword followed by a semicolon separated list.
func (s *Statement) ForFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "for",
		open:      "for ",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Switch renders the keyword followed by a semicolon separated list.
func Switch(conditions ...Code) *Statement {
	return newStatement().Switch(conditions...)
}

// Switch renders the keyword followed by a semicolon separated list.
func (g *Group) Switch(conditions ...Code) *Statement {
	s := Switch(conditions...)
	g.items = append(g.items, s)
	return s
}

// Switch renders the keyword followed by a semicolon separated list.
func (s *Statement) Switch(conditions ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     conditions,
		multi:     false,
		name:      "switch",
		open:      "switch ",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// SwitchFunc renders the keyword followed by a semicolon separated list.
func SwitchFunc(f func(*Group)) *Statement {
	return newStatement().SwitchFunc(f)
}

// SwitchFunc renders the keyword followed by a semicolon separated list.
func (g *Group) SwitchFunc(f func(*Group)) *Statement {
	s := SwitchFunc(f)
	g.items = append(g.items, s)
	return s
}

// SwitchFunc renders the keyword followed by a semicolon separated list.
func (s *Statement) SwitchFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "switch",
		open:      "switch ",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Case renders the keyword followed by a comma separated list.
func Case(cases ...Code) *Statement {
	return newStatement().Case(cases...)
}

// Case renders the keyword followed by a comma separated list.
func (g *Group) Case(cases ...Code) *Statement {
	s := Case(cases...)
	g.items = append(g.items, s)
	return s
}

// Case renders the keyword followed by a comma separated list.
func (s *Statement) Case(cases ...Code) *Statement {
	g := &Group{
		close:     ":",
		items:     cases,
		multi:     false,
		name:      "case",
		open:      "case ",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// CaseFunc renders the keyword followed by a comma separated list.
func CaseFunc(f func(*Group)) *Statement {
	return newStatement().CaseFunc(f)
}

// CaseFunc renders the keyword followed by a comma separated list.
func (g *Group) CaseFunc(f func(*Group)) *Statement {
	s := CaseFunc(f)
	g.items = append(g.items, s)
	return s
}

// CaseFunc renders the keyword followed by a comma separated list.
func (s *Statement) CaseFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ":",
		multi:     false,
		name:      "case",
		open:      "case ",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Len renders the len built-in function.
func Len(v Code) *Statement {
	return newStatement().Len(v)
}

// Len renders the len built-in function.
func (g *Group) Len(v Code) *Statement {
	s := Len(v)
	g.items = append(g.items, s)
	return s
}

// Len renders the len built-in function.
func (s *Statement) Len(v Code) *Statement {
	g := &Group{
		close:     ".length",
		items:     []Code{v},
		multi:     false,
		name:      "len",
		open:      "",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Print renders the print built-in function.
func Print(args ...Code) *Statement {
	return newStatement().Print(args...)
}

// Print renders the print built-in function.
func (g *Group) Print(args ...Code) *Statement {
	s := Print(args...)
	g.items = append(g.items, s)
	return s
}

// Print renders the print built-in function.
func (s *Statement) Print(args ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     args,
		multi:     false,
		name:      "print",
		open:      "console.log(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// PrintFunc renders the print built-in function.
func PrintFunc(f func(*Group)) *Statement {
	return newStatement().PrintFunc(f)
}

// PrintFunc renders the print built-in function.
func (g *Group) PrintFunc(f func(*Group)) *Statement {
	s := PrintFunc(f)
	g.items = append(g.items, s)
	return s
}

// PrintFunc renders the print built-in function.
func (s *Statement) PrintFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		multi:     false,
		name:      "print",
		open:      "console.log(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Types renders a comma separated list enclosed by square brackets. Use for type parameters and constraints.
func Types(types ...Code) *Statement {
	return newStatement().Types(types...)
}

// Types renders a comma separated list enclosed by square brackets. Use for type parameters and constraints.
func (g *Group) Types(types ...Code) *Statement {
	s := Types(types...)
	g.items = append(g.items, s)
	return s
}

// Types renders a comma separated list enclosed by square brackets. Use for type parameters and constraints.
func (s *Statement) Types(types ...Code) *Statement {
	g := &Group{
		close:     ">",
		items:     types,
		multi:     false,
		name:      "types",
		open:      "<",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// TypesFunc renders a comma separated list enclosed by square brackets. Use for type parameters and constraints.
func TypesFunc(f func(*Group)) *Statement {
	return newStatement().TypesFunc(f)
}

// TypesFunc renders a comma separated list enclosed by square brackets. Use for type parameters and constraints.
func (g *Group) TypesFunc(f func(*Group)) *Statement {
	s := TypesFunc(f)
	g.items = append(g.items, s)
	return s
}

// TypesFunc renders a comma separated list enclosed by square brackets. Use for type parameters and constraints.
func (s *Statement) TypesFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ">",
		multi:     false,
		name:      "types",
		open:      "<",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Union renders a pipe separated list. Use for union type constraints.
func Union(types ...Code) *Statement {
	return newStatement().Union(types...)
}

// Union renders a pipe separated list. Use for union type constraints.
func (g *Group) Union(types ...Code) *Statement {
	s := Union(types...)
	g.items = append(g.items, s)
	return s
}

// Union renders a pipe separated list. Use for union type constraints.
func (s *Statement) Union(types ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     types,
		multi:     false,
		name:      "union",
		open:      "",
		separator: "|",
	}
	*s = append(*s, g)
	return s
}

// UnionFunc renders a pipe separated list. Use for union type constraints.
func UnionFunc(f func(*Group)) *Statement {
	return newStatement().UnionFunc(f)
}

// UnionFunc renders a pipe separated list. Use for union type constraints.
func (g *Group) UnionFunc(f func(*Group)) *Statement {
	s := UnionFunc(f)
	g.items = append(g.items, s)
	return s
}

// UnionFunc renders a pipe separated list. Use for union type constraints.
func (s *Statement) UnionFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		multi:     false,
		name:      "union",
		open:      "",
		separator: "|",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Boolean renders the boolean identifier.
func Boolean() *Statement {
	return newStatement().Boolean()
}

// Boolean renders the boolean identifier.
func (g *Group) Boolean() *Statement {
	s := Boolean()
	g.items = append(g.items, s)
	return s
}

// Boolean renders the boolean identifier.
func (s *Statement) Boolean() *Statement {
	t := token{
		content: "boolean",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Number renders the number identifier.
func Number() *Statement {
	// notest
	return newStatement().Number()
}

// Number renders the number identifier.
func (g *Group) Number() *Statement {
	// notest
	s := Number()
	g.items = append(g.items, s)
	return s
}

// Number renders the number identifier.
func (s *Statement) Number() *Statement {
	// notest
	t := token{
		content: "number",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// String renders the string identifier.
func String() *Statement {
	// notest
	return newStatement().String()
}

// String renders the string identifier.
func (g *Group) String() *Statement {
	// notest
	s := String()
	g.items = append(g.items, s)
	return s
}

// String renders the string identifier.
func (s *Statement) String() *Statement {
	// notest
	t := token{
		content: "string",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Bigint renders the bigint identifier.
func Bigint() *Statement {
	// notest
	return newStatement().Bigint()
}

// Bigint renders the bigint identifier.
func (g *Group) Bigint() *Statement {
	// notest
	s := Bigint()
	g.items = append(g.items, s)
	return s
}

// Bigint renders the bigint identifier.
func (s *Statement) Bigint() *Statement {
	// notest
	t := token{
		content: "bigint",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Any renders the any identifier.
func Any() *Statement {
	// notest
	return newStatement().Any()
}

// Any renders the any identifier.
func (g *Group) Any() *Statement {
	// notest
	s := Any()
	g.items = append(g.items, s)
	return s
}

// Any renders the any identifier.
func (s *Statement) Any() *Statement {
	// notest
	t := token{
		content: "any",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Undefined renders the undefined identifier.
func Undefined() *Statement {
	// notest
	return newStatement().Undefined()
}

// Undefined renders the undefined identifier.
func (g *Group) Undefined() *Statement {
	// notest
	s := Undefined()
	g.items = append(g.items, s)
	return s
}

// Undefined renders the undefined identifier.
func (s *Statement) Undefined() *Statement {
	// notest
	t := token{
		content: "undefined",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Break renders the break keyword.
func Break() *Statement {
	// notest
	return newStatement().Break()
}

// Break renders the break keyword.
func (g *Group) Break() *Statement {
	// notest
	s := Break()
	g.items = append(g.items, s)
	return s
}

// Break renders the break keyword.
func (s *Statement) Break() *Statement {
	// notest
	t := token{
		content: "break",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Default renders the default keyword.
func Default() *Statement {
	// notest
	return newStatement().Default()
}

// Default renders the default keyword.
func (g *Group) Default() *Statement {
	// notest
	s := Default()
	g.items = append(g.items, s)
	return s
}

// Default renders the default keyword.
func (s *Statement) Default() *Statement {
	// notest
	t := token{
		content: "default",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Function renders the function keyword.
func Function() *Statement {
	// notest
	return newStatement().Function()
}

// Function renders the function keyword.
func (g *Group) Function() *Statement {
	// notest
	s := Function()
	g.items = append(g.items, s)
	return s
}

// Function renders the function keyword.
func (s *Statement) Function() *Statement {
	// notest
	t := token{
		content: "function",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Else renders the else keyword.
func Else() *Statement {
	// notest
	return newStatement().Else()
}

// Else renders the else keyword.
func (g *Group) Else() *Statement {
	// notest
	s := Else()
	g.items = append(g.items, s)
	return s
}

// Else renders the else keyword.
func (s *Statement) Else() *Statement {
	// notest
	t := token{
		content: "else",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Export renders the export keyword.
func Export() *Statement {
	// notest
	return newStatement().Export()
}

// Export renders the export keyword.
func (g *Group) Export() *Statement {
	// notest
	s := Export()
	g.items = append(g.items, s)
	return s
}

// Export renders the export keyword.
func (s *Statement) Export() *Statement {
	// notest
	t := token{
		content: "export",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Import renders the import keyword.
func Import() *Statement {
	// notest
	return newStatement().Import()
}

// Import renders the import keyword.
func (g *Group) Import() *Statement {
	// notest
	s := Import()
	g.items = append(g.items, s)
	return s
}

// Import renders the import keyword.
func (s *Statement) Import() *Statement {
	// notest
	t := token{
		content: "import",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Const renders the const keyword.
func Const() *Statement {
	// notest
	return newStatement().Const()
}

// Const renders the const keyword.
func (g *Group) Const() *Statement {
	// notest
	s := Const()
	g.items = append(g.items, s)
	return s
}

// Const renders the const keyword.
func (s *Statement) Const() *Statement {
	// notest
	t := token{
		content: "const",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Class renders the class keyword.
func Class() *Statement {
	// notest
	return newStatement().Class()
}

// Class renders the class keyword.
func (g *Group) Class() *Statement {
	// notest
	s := Class()
	g.items = append(g.items, s)
	return s
}

// Class renders the class keyword.
func (s *Statement) Class() *Statement {
	// notest
	t := token{
		content: "class",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Type renders the type keyword.
func Type() *Statement {
	// notest
	return newStatement().Type()
}

// Type renders the type keyword.
func (g *Group) Type() *Statement {
	// notest
	s := Type()
	g.items = append(g.items, s)
	return s
}

// Type renders the type keyword.
func (s *Statement) Type() *Statement {
	// notest
	t := token{
		content: "type",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Continue renders the continue keyword.
func Continue() *Statement {
	// notest
	return newStatement().Continue()
}

// Continue renders the continue keyword.
func (g *Group) Continue() *Statement {
	// notest
	s := Continue()
	g.items = append(g.items, s)
	return s
}

// Continue renders the continue keyword.
func (s *Statement) Continue() *Statement {
	// notest
	t := token{
		content: "continue",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Var renders the var keyword.
func Var() *Statement {
	// notest
	return newStatement().Var()
}

// Var renders the var keyword.
func (g *Group) Var() *Statement {
	// notest
	s := Var()
	g.items = append(g.items, s)
	return s
}

// Var renders the var keyword.
func (s *Statement) Var() *Statement {
	// notest
	t := token{
		content: "var",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// New renders the new keyword.
func New() *Statement {
	// notest
	return newStatement().New()
}

// New renders the new keyword.
func (g *Group) New() *Statement {
	// notest
	s := New()
	g.items = append(g.items, s)
	return s
}

// New renders the new keyword.
func (s *Statement) New() *Statement {
	// notest
	t := token{
		content: "new",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Interface renders the interface keyword.
func Interface() *Statement {
	// notest
	return newStatement().Interface()
}

// Interface renders the interface keyword.
func (g *Group) Interface() *Statement {
	// notest
	s := Interface()
	g.items = append(g.items, s)
	return s
}

// Interface renders the interface keyword.
func (s *Statement) Interface() *Statement {
	// notest
	t := token{
		content: "interface",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Namespace renders the namespace keyword.
func Namespace() *Statement {
	// notest
	return newStatement().Namespace()
}

// Namespace renders the namespace keyword.
func (g *Group) Namespace() *Statement {
	// notest
	s := Namespace()
	g.items = append(g.items, s)
	return s
}

// Namespace renders the namespace keyword.
func (s *Statement) Namespace() *Statement {
	// notest
	t := token{
		content: "namespace",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}
