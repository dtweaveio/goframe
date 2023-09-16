package main

import (
	"dtweave.io/goframe/app"
	"dtweave.io/goframe/samples/myoptions"
)

func main() {
	opts := myoptions.NewMyOptions()

	builder := app.NewAppBuilder()
	app := builder.
		WithName("myapp").
		WithShortName("myapp").
		WithDesc("this is my app").
		WithOptions(opts).
		WithRun(func(basename string) error {
			return nil
		}).
		Build()

	app.Run()
}
