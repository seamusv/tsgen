package tsgen

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFile_RenderHeaders(t *testing.T) {
	tests := []struct {
		name   string
		file   func() *File
		output string
	}{
		{
			name: "single header",
			file: func() *File {
				f := NewFile()
				f.HeaderComment("this file is generated")
				return f
			},
			output: "// this file is generated\n",
		},
		{
			name: "multiple headers",
			file: func() *File {
				f := NewFile()
				f.HeaderComment("this file is generated")
				f.HeaderComment("Welcome to the Jungle")
				return f
			},
			output: "// this file is generated\n// Welcome to the Jungle\n",
		},
		{
			name: "simple clas1s",
			file: func() *File {
				f := NewFile(WithPrettierFormatter())

				f.Import("@protobuf-ts/grpcweb-transport", "GrpcWebFetchTransport")
				f.Import("@protobuf-ts/runtime-rpc", "RpcOptions")
				f.Import("@protobuf-ts/runtime-rpc", "ServerStreamingCall")
				f.Import("@protobuf-ts/runtime-rpc", "UnaryCall")
				f.ImportDefault("@sparx/protobuf/sparks_pb", "pb")
				f.ImportDefault("@sparx/protobuf/sparks.client", "pb_client")

				f.Const().Id("transport").Op("=").New().Id("GrpcWebFetchTransport").
					Call(Values(Dict{
						Id("baseUrl"): Id("window.location.origin"),
						Id("format"):  Lit("binary"),
					})).
					Semicolon().
					Line()

				services := []string{"AccountService", "AuthService", "GameService", "UserService"}
				for _, service := range services {
					f.Const().Id(strcase.ToLowerCamel(service) + "Client").Op("=").
						New().Id("pb_client").Dot(service + "Client").
						Call(Id("transport")).
						Semicolon()
				}

				f.Line()

				for _, service := range services {
					f.Type().Id(service).Op("=").Values(Dict{
						Id("bar"): Do(func(s *Statement) {
							s.Id("UnaryRequest").Types(Id("pb").Dot("GetAccountRequest"), Id("pb").Dot("GetAccountResponse"))
						}),
						Id("baz"): Do(func(s *Statement) {
							s.Id("UnaryRequest").Types(Id("pb").Dot("GetAccountRequest"), Id("pb").Dot("GetAccountResponse"))
						}),
					})
				}

				f.Line()

				f.Export().Const().Id("nop").Op("=").Params().Op("=>").Block(Return(Undefined()))

				f.Line()

				f.Export().Namespace().Id("StudioGrpc").BlockFunc(func(bg *Group) {
					bg.Export().Interface().Id("Services").ValuesFunc(func(sg *Group) {
						dict := Dict{}
						for _, service := range services {
							dict[Id(strings.ReplaceAll(service, "Service", ""))] = Id(service)
						}
						sg.Add(dict)
					})

					bg.Line()

					bg.Export().Const().Id("Wired").Op(":").Id("Services").Op("=").ValuesFunc(func(vg *Group) {
						servicesDict := Dict{}
						for _, service := range services {
							servicesDict[Id(strings.ReplaceAll(service, "Service", ""))] =
								Params(Dict{Id("r"): Any(), Id("c"): Any()}).Op("=>").Id("asUnary").Call(Id("r"), Id("c"))
						}
						vg.Add(servicesDict)
					})
				})

				fmt.Println("-------")
				fmt.Println(f.GoString())
				fmt.Println("-------")
				return f
			},
			output: "",
		},
		{
			name: "simple class",
			file: func() *File {
				f := NewFile(WithPrettierFormatter())

				f.ImportDefault("react", "React")

				f.Interface().Id("Props").Values(Dict{
					Id("name"): Id("string"),
				})

				f.Line()

				f.Export().Const().Id("MyComponent").Op("=").Params(Dict{Id("props"): Id("Props")}).Op("=>").Block(
					Return(Id("<div>Welcome to the Jungle, {props.name}!</div>")),
				)

				fmt.Println("-------")
				fmt.Printf("%#v", f)
				fmt.Println("-------")
				return f
			},
			output: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.file().GoString()
			assert.Equal(t, tt.output, output)
		})
	}
}

func TestFile_RenderImport(t *testing.T) {
	tests := []struct {
		name   string
		file   func() *File
		output string
	}{
		{
			name: "single import",
			file: func() *File {
				f := NewFile()
				f.Import("@sparx/protobuf", "AssetItem")
				return f
			},
			output: "import { AssetItem } from '@sparx/protobuf';\n",
		},
		{
			name: "default import",
			file: func() *File {
				f := NewFile()
				f.ImportDefault("@sparx/protobuf", "pb")
				return f
			},
			output: "import * as pb from '@sparx/protobuf';\n",
		},
		{
			name: "import multiple from module",
			file: func() *File {
				f := NewFile()
				f.Import("@sparx/protobuf", "AssetItem")
				f.Import("@sparx/protobuf", "AssetDriver")
				return f
			},
			output: "import { AssetDriver, AssetItem } from '@sparx/protobuf';\n",
		},
		{
			name: "import multiple with aliases from module",
			file: func() *File {
				f := NewFile()
				f.Import("@sparx/protobuf", "AssetItem")
				f.Import("@sparx/protobuf", "AssetDriver")
				f.ImportAlias("@sparx/protobuf", "Image", "Img")
				return f
			},
			output: "import { AssetDriver, AssetItem, Image as Img } from '@sparx/protobuf';\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := tt.file().GoString()
			assert.Equal(t, tt.output, output)
		})
	}
}
