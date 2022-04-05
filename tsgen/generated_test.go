package tsgen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExport(t *testing.T) {
	tests := []struct {
		name string
		op   func() *Statement
		want string
	}{
		{
			name: "simple export",
			op: func() *Statement {
				return Export()
			},
			want: "export",
		},
		{
			name: "export const",
			op: func() *Statement {
				return Export().Const().Id("foo").Op("=").Lit(42)
			},
			want: "export const foo = 42",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.op().TypescriptString(), "Export()")
		})
	}
}
