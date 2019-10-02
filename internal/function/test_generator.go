package function

import (
"fmt"
"github.com/src-d/go-mysql-server/sql"
)

type TestGenerator struct {
}

func NewTestGenerator() (*TestGenerator, error) {
	return &TestGenerator{}, nil
}

func (g *TestGenerator) Next() (interface{}, error) {
	return 1, nil
}

func (g *TestGenerator) Close() error {
	return nil
}

var _ sql.Generator = (*TestGenerator)(nil)

type TestGenerate struct {}

func NewTestGenerate() sql.Expression {
	return &TestGenerate{}
}

func (b *TestGenerate) String() string {
	return fmt.Sprintf("testgenerate")
}

// Type implements the sql.Expression interface
func (*TestGenerate) Type() sql.Type {
	return sql.Array(sql.Int8)
}

func (b *TestGenerate) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	return NewTestGenerate(), nil
}

// Children implements the Expression interface.
func (b *TestGenerate) Children() []sql.Expression {
	return []sql.Expression{}
}

// IsNullable implements the Expression interface.
func (*TestGenerate) IsNullable() bool {
	return false
}

// Resolved implements the Expression interface.
func (b *TestGenerate) Resolved() bool {
	return true
}

// Eval implements the sql.Expression interface.
func (b *TestGenerate) Eval(ctx *sql.Context, row sql.Row) (interface{}, error) {
	span, ctx := ctx.Span("gitbase.TestGenerate")
	defer span.Finish()

	bg, err := NewTestGenerator()
	if err != nil {
		return nil, err
	}

	return bg, nil
}

