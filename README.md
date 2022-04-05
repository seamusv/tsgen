# TSGen

Typescript generator for Go.

Inspired and copied from Dave Brophy's [Jennifer](https://github.com/dave/jennifer) project.

```go
package main

import (
    "fmt"

    . "github.com/seamusv/tsgen/gen"
)

func main() {
	f := NewFile(WithPrettierFormatter())

	f.ImportDefault("react", "React")

	f.Interface().Id("Props").Values(Dict{
		Id("name"): Id("string"),
	})

	f.Line()

	f.Export().Const().Id("MyComponent").Op("=").Params(Dict{Id("props"): Id("Props")}).Op("=>").Block(
		Return(Id("<div>Welcome to the Jungle, {props.name}!</div>")),
	)
	fmt.Printf("%#v", f)
}
```
Output:
```typescript
import * as React from 'react';

interface Props {
    name: string;
}

export const MyComponent = (props: Props) => {
    return <div>Welcome to the Jungle, {props.name}!</div>;
};
```

### Install
```
go get -u github.com/seamusv/tsgen/tsgen
```

## WIP