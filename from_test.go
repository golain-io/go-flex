package goflex

import (
	"fmt"
	"testing"
)

func TestFrom(t *testing.T) {
	B := FromBuilder{}
	B.From("test").ToFlux()
	fmt.Println(B)
}
