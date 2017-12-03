package syntax

import "bytes"

// A Query contains parse tree of a query.
type Query struct {
	Input string  // the original input query string
	Expr  []*Expr // expressions in this query
}

// An Expr describes an expression in a query.
type Expr struct {
	Pos       int       // the starting character position of the query expression
	Not       bool      // the expression is negated (e.g., -term or -field:term)
	Field     string    // the field that this expression applies to
	Value     string    // the raw field value
	ValueType TokenType // the type of the value
}

func (e Expr) String() string {
	var buf bytes.Buffer
	if e.Not {
		buf.WriteByte('-')
	}
	if e.Field != "" {
		buf.WriteString(e.Field)
		buf.WriteByte(':')
	}
	if e.ValueType == TokenPattern {
		buf.WriteByte('/')
	}
	buf.WriteString(e.Value)
	if e.ValueType == TokenPattern {
		buf.WriteByte('/')
	}
	return buf.String()
}
