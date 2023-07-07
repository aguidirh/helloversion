package helloversioning

import v2 "github.com/aguidirh/helloversioning/v2"

// version simulates a flag to forward a workflow to another module version
var version = "v1"

func Hello() string {
	if version == "v1" {
		return "Hello, world v1."
	} else {
		return v2.Hello()
	}
}
