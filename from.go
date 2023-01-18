package goflex

import (
	"bytes"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/lann/builder"
)

// from() funtion representation. this is the starting point for most FluxQL queries
type fromQuery struct {
	influxOptions influxdb2.Options // client options for influxdb-go-client
	client        influxdb2.Client  // actual instantiated influxdb client
	From          string            // holds the value of the bucket to query
	Range         string            // stringified range arguments (in fluxql format)
	Filter        *FilterBuilder    // Filter Builder
	Sum           Fluxer            // sample math operation
}

func (f *fromQuery) ToFlux() (qString string, args []interface{}, err error) {
	// cannot build query without a bucket name
	if f.From == "" {
		return "", nil, fmt.Errorf("from() function requires a bucket name")
	}
	// need range to query
	if f.Range == "" {
		return "", nil, fmt.Errorf("from() function requires a start range")
	}

	flux := &bytes.Buffer{}

	// start the from function
	flux.WriteString("from(bucket: \"")
	flux.WriteString(f.From)
	flux.WriteString("\")")

	// add the range
	flux.WriteString(" |> range(")
	flux.WriteString(f.Range)
	flux.WriteString(")")

	if f.Filter != nil {
		filter := builder.GetStruct(*f.Filter).(filterData)
		err = addFilter(filter, flux)
	}
	qString = flux.String()
	return

}

// FromBuilder is the builder type for the from('bucket') function
// creates the actual fromQuery variable
type FromBuilder builder.Builder

// initialise the builder and return an empty Builder
func init() {
	builder.Register(FromBuilder{}, fromQuery{})
}

func (b FromBuilder) From(from string) FromBuilder {
	return builder.Set(b, "From", from).(FromBuilder)
}

func (b FromBuilder) Range(t map[string]string) FromBuilder {
	if t["stop"] == "" {
		return builder.Set(b, "Range", fmt.Sprintf("start:%s", t["start"])).(FromBuilder)
	}
	return builder.Set(b, "Range", fmt.Sprintf("start:%s, stop:%s", t["start"], t["stop"])).(FromBuilder)
}

// generates string query to pass to influx client
func (b FromBuilder) ToFlux() (qString string, args []interface{}, err error) {
	data := builder.GetStruct(b).(fromQuery)
	return data.ToFlux()
}

// func (b FromBuilder) Query() (*influxapi.QueryTableResult, error) {
// 	if f.client == nil {
// 		return nil, ClientNotSet
// 	}
// 	queryAPI := f.client.QueryAPI(f.org)
// 	return QueryWith(&queryAPI, f)
// }
