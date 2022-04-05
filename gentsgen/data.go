package main

var keywords = []string{"break", "default", "function", "else", "export", "const", "class", "type", "continue", "var", "new", "interface", "namespace"}

var identifiers = []string{"boolean", "number", "string", "bigint", "any", "undefined"}

var groups = []struct {
	name        string   // name of the function / method
	comment     string   // comment appended to name
	variadic    bool     // is the parameter variadic?
	opening     string   // opening token
	closing     string   // closing token
	separator   string   // separator token
	multi       bool     // items are always on multiple lines
	parameters  []string // parameter names
	preventFunc bool     // prevent the fooFunc function/method
	preventSemi bool     // prevent the semi-colon
}{
	{
		name:      "Semicolon",
		comment:   "renders a semicolon",
		opening:   "",
		closing:   ";",
		separator: "",
	},
	{
		name:       "Parens",
		comment:    "renders a single item in parenthesis. Use for type conversion or to specify evaluation order.",
		variadic:   false,
		opening:    "(",
		closing:    ")",
		separator:  "",
		parameters: []string{"item"},
	},
	{
		name:       "List",
		comment:    "renders a comma separated list. Use for multiple return functions.",
		variadic:   true,
		opening:    "",
		closing:    "",
		separator:  ",",
		parameters: []string{"items"},
	},
	{
		name:       "Values",
		comment:    "renders a comma separated list enclosed by curly braces. Use for slice or composite literals.",
		variadic:   true,
		opening:    "{",
		closing:    "}",
		separator:  ",",
		parameters: []string{"values"},
	},
	{
		name:       "Method",
		comment:    "renders a method call. Use for method calls.",
		opening:    "",
		closing:    "",
		separator:  "",
		parameters: []string{"name"},
	},
	{
		name:       "Index",
		comment:    "renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.",
		variadic:   true,
		opening:    "[",
		closing:    "]",
		separator:  ":",
		parameters: []string{"items"},
	},
	{
		name:       "Block",
		comment:    "renders a statement list enclosed by curly braces. Use for code blocks. A special case applies when used directly after Case or Default, where the braces are omitted. This allows use in switch and select statements.",
		variadic:   true,
		opening:    "{",
		closing:    "}",
		multi:      true,
		parameters: []string{"statements"},
	},
	{
		name:       "Defs",
		comment:    "renders a statement list enclosed in parenthesis. Use for definition lists.",
		variadic:   true,
		opening:    "(",
		closing:    ")",
		multi:      true,
		parameters: []string{"definitions"},
	},
	{
		name:       "Call",
		comment:    "renders a comma separated list enclosed by parenthesis. Use for function calls.",
		variadic:   true,
		opening:    "(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"params"},
	},
	{
		name:       "Params",
		comment:    "renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.",
		variadic:   true,
		opening:    "(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"params"},
	},
	{
		name:       "Assert",
		comment:    "renders a period followed by a single item enclosed by parenthesis. Use for type assertions.",
		variadic:   false,
		opening:    ".(",
		closing:    ")",
		separator:  "",
		parameters: []string{"typ"},
	},
	{
		name:       "If",
		comment:    "renders the keyword followed by a semicolon separated list.",
		variadic:   true,
		opening:    "if ",
		closing:    "",
		separator:  ";",
		parameters: []string{"conditions"},
	},
	{
		name:       "Return",
		comment:    "renders the keyword followed by a comma separated list.",
		variadic:   false,
		opening:    "return ",
		closing:    "",
		separator:  ",",
		parameters: []string{"result"},
	},
	{
		name:       "For",
		comment:    "renders the keyword followed by a semicolon separated list.",
		variadic:   true,
		opening:    "for ",
		closing:    "",
		separator:  ";",
		parameters: []string{"conditions"},
	},
	{
		name:       "Switch",
		comment:    "renders the keyword followed by a semicolon separated list.",
		variadic:   true,
		opening:    "switch ",
		closing:    "",
		separator:  ";",
		parameters: []string{"conditions"},
	},
	{
		name:       "Case",
		comment:    "renders the keyword followed by a comma separated list.",
		variadic:   true,
		opening:    "case ",
		closing:    ":",
		separator:  ",",
		parameters: []string{"cases"},
	},
	{
		name:       "Len",
		comment:    "renders the len built-in function.",
		variadic:   false,
		opening:    "",
		closing:    ".length",
		separator:  ",",
		parameters: []string{"v"},
	},
	{
		name:       "Print",
		comment:    "renders the print built-in function.",
		variadic:   true,
		opening:    "console.log(",
		closing:    ")",
		separator:  ",",
		parameters: []string{"args"},
	},
	{
		name:       "Types",
		comment:    "renders a comma separated list enclosed by square brackets. Use for type parameters and constraints.",
		variadic:   true,
		opening:    "<",
		closing:    ">",
		separator:  ",",
		parameters: []string{"types"},
	},
	{
		name:       "Union",
		comment:    "renders a pipe separated list. Use for union type constraints.",
		variadic:   true,
		opening:    "",
		closing:    "",
		separator:  "|",
		parameters: []string{"types"},
	},
}
