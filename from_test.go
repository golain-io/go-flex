package goflex

import (
	"fmt"
	"testing"
)

func TestFrom(t *testing.T) {
	f, _, err := From("g-client1").Range(map[string]string{"start": "-30d", "stop": "-1s"}).ToFlux()
	fmt.Println(f)
	fmt.Println(err)
}
