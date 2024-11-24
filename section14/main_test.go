package main

import (
	"testing"
)

var testAppName = "GOAPP"
var testVersion = "1.0"

var Debug bool = true

func TestIsOne(t *testing.T) {
	v := IsOne(1)
	if v != true {
		t.Errorf("IsOne(1) = %v; want true", v)
	}

	v = IsOne(2)
	if v != false {
		t.Errorf("IsOne(2) = %v; want false", v)
	}
}
