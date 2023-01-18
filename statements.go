package goflex

import "github.com/lann/builder"

type StatementBuilderType builder.Builder

func (b StatementBuilderType) From(from string) FromBuilder {
	return FromBuilder(b).From(from)
}

// StatementBuilder is a parent builder for other builders, e.g. SelectBuilder.
var StatementBuilder = StatementBuilderType(builder.EmptyBuilder)

// From returns a new FromBuilder with the given bucket name.
func From(bucket string) FromBuilder {
	return StatementBuilder.From(bucket)
}
