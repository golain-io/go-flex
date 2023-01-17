package goflex

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/lann/builder"
)

// from() funtion representation. this is the starting point for most FluxQL queries
type fromQuery struct {
	influxOptions influxdb2.Options // client options for influxdb-go-client
	client        influxdb2.Client  // actual instantiated influxdb client
	From          string            // holds the value of the bucket to query
	Range         Fluxer            // range to query
	Filter        *FilterBuilder    // Filter Builder
	Sum           Fluxer            // sample math operation
}

func (f *fromQuery) ToSql() (qString string, args []interface{}, err error) {
	// cannot build query without a bucket name
	if f.From == "" {
		return "", nil, fmt.Errorf("from() function requires a bucket name")
	}
	// need range to query
	if f.Range == nil {
		return "", nil, fmt.Errorf("from() function requires a range")
	}

	return qString, nil, nil
}

// FromBuilder is the builder type for the from('bucket') function
// creates the actual fromQuery variable
type FromBuilder builder.Builder

// initialise the builder and return an empty Builder
func init() {
	builder.Register(FromBuilder{}, FromBuilder{})
}

func (b FromBuilder) From(from string) FromBuilder {
	return builder.Set(b, "From", from).(FromBuilder)
}

func (b FromBuilder) Range(time ...string) FromBuilder {
	return builder.Extend(b, "Range", time).(FromBuilder)
}

func (b FromBuilder) Filter(time ...string) FromBuilder {
	return builder.Extend(b, "Range", time).(FromBuilder)
}

// generates string query to pass to influx client
func (b FromBuilder) ToFlux() (qString string, args []interface{}, err error) {
	data := builder.GetStruct(b).(fromQuery)
	return data.ToSql()
}

// func (b FromBuilder) Query() (*influxapi.QueryTableResult, error) {
// 	if f.client == nil {
// 		return nil, ClientNotSet
// 	}
// 	queryAPI := f.client.QueryAPI(f.org)
// 	return QueryWith(&queryAPI, f)
// }
