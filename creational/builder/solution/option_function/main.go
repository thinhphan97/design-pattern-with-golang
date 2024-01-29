package main

import "github.com/thinhphan97/design-pattern-with-golang/creational/builder/solution/option_function/internal"

func main() {
	service := internal.NewService(
		internal.WithName("Complex Service"),
		internal.WithStdLogger(),
		internal.WithEmailNotifier(),
		internal.WithMySQLDataLayer(),
	)

	service.DoBusiness()
}
