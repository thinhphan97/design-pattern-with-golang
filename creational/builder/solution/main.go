package main

import (
	"github.com/thinhphan97/design-pattern-with-golang/creational/builder/solution/internal"
)

func main() {
	director := internal.NewDirector()
	builder := internal.NewBuilder()

	service := director.BuildService(builder)
	service.DoBusiness()
}
