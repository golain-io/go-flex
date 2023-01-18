package goflex

import (
	"context"
	"fmt"
	"io"

	influxapi "github.com/influxdata/influxdb-client-go/v2/api"
)

// client not set error is reused in multiple different places
var ErrClientNotSet = fmt.Errorf("cannot run; no Client set (Query)")

// Fluxer is the interface that wraps the ToFlux method.
//
// ToFlux will return a FluxQL(string) representation of the Fluxer (with arguments)
// It may also return an error.
//
// For variables within Builder types, ToFlux is consumed by the appendToFlux method
// which will append the result of ToFlux to the query string
//
// []interface{} is used to pass arguments to the query string
type Fluxer interface {
	ToFlux() (string, []interface{}, error)
}

// universal function to actually pass the generated query to the influxdb client
// results QueryTableResult as with the influx db client package
func QueryWith(api *influxapi.QueryAPI, f Fluxer) (res *influxapi.QueryTableResult, err error) {
	// @TODO
	query, err := "", nil
	if err != nil {
		return
	}
	return (*api).Query(context.Background(), query)
}

// takes a filterData struct and appends it to the passed query string
func addFilter(f filterData, w io.Writer) error {
	// @TODO
	return nil
}
