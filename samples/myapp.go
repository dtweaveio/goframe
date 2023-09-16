package main

import (
	"github.com/dtweaveio/goframe/app"
	"github.com/dtweaveio/goframe/samples/myoptions"
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
