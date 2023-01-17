package goflex

import "github.com/lann/builder"

// filter builder type
type FilterBuilder builder.Builder

type filterData struct {
	Field       Fluxer // field parameter
	Tag         Fluxer // by tag
	Measurement Fluxer // by measurement name
	// Conditionals
	Eq            string
	NotEq         string
	LessThan      string
	GreaterThan   string
	LessThanEq    string
	GreaterThanEq string
	// Conjunctive
	And bool
	Or  bool
}

// initialise the builder and return an empty Builder
func init() {
	builder.Register(FilterBuilder{}, filterData{})
}

func (b FilterBuilder) Field(field string) FilterBuilder {
	return builder.Set(b, "Field", field).(FilterBuilder)
}

func (b FilterBuilder) Tag(tag string) FilterBuilder {
	return builder.Set(b, "Tag", tag).(FilterBuilder)
}

func (b FilterBuilder) Measurement(measurement string) FilterBuilder {
	return builder.Set(b, "Measurement", measurement).(FilterBuilder)
}

func (b FilterBuilder) Eq(eq string) FilterBuilder {
	return builder.Set(b, "Eq", eq).(FilterBuilder)
}

func (b FilterBuilder) NotEq(neq string) FilterBuilder {
	return builder.Set(b, "NotEq", neq).(FilterBuilder)
}

func (b FilterBuilder) LessThan(lt string) FilterBuilder {
	return builder.Set(b, "LessThan", lt).(FilterBuilder)
}

func (b FilterBuilder) GreaterThan(gt string) FilterBuilder {
	return builder.Set(b, "GreaterThan", gt).(FilterBuilder)
}

func (b FilterBuilder) LessThanEq(lte string) FilterBuilder {
	return builder.Set(b, "LessThanEq", lte).(FilterBuilder)
}

func (b FilterBuilder) GreaterThanEq(gte string) FilterBuilder {
	return builder.Set(b, "GreaterThanEq", gte).(FilterBuilder)
}

func (b FilterBuilder) And(and string) FilterBuilder {
	return builder.Set(b, "And", and).(FilterBuilder)
}

func (b FilterBuilder) Or(or string) FilterBuilder {
	return builder.Set(b, "Or", or).(FilterBuilder)
}

// generates string query to pass to influx client
func (b FilterBuilder) ToFlux() (qString string, args []interface{}, err error) {
	return
}
