package main

import (
	"dtweave.io/goframe/app"
)

func main() {
	builder := app.NewAppBuilder()
	app := builder.
		WithName("dataweave").
		WithShortName("myapp").
		WithDesc("des").
		WithNoConfig(true).
		WithRun(func(basename string) error {
			return nil
		}).
		Build()

	app.Run()
}
