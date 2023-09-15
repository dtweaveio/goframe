package app

import (
	"github.com/spf13/cobra"
)

// App is the main structure of a cli application.
// It is recommended that an app be created with the app_boot.NewApp() function.
type App struct {
	name        string
	shortName   string
	description string
	preRunFunc  RunFunc
	runFunc     RunFunc
	postRunFunc RunFunc
	noVersion   bool
	noConfig    bool
	options     Options
	cmd         *cobra.Command
	commands    []*Command
	args        cobra.PositionalArgs
}

// RunFunc defines the application's startup callback function.
type RunFunc func(basename string) error

type AppBuilder struct {
	app App
}

func NewAppBuilder() *AppBuilder {
	return &AppBuilder{}
}

func (b *AppBuilder) Build() App {
	return b.app
}

func (b *AppBuilder) WithName(name string) *AppBuilder {
	b.app.name = name
	return b
}

func (b *AppBuilder) WithShortName(shortName string) *AppBuilder {
	b.app.shortName = shortName
	return b
}

func (b *AppBuilder) WithDesc(desc string) *AppBuilder {
	b.app.description = desc
	return b
}

func (b *AppBuilder) WithPreRun(preRun RunFunc) *AppBuilder {
	b.app.preRunFunc = preRun
	return b
}

func (b *AppBuilder) WithRun(run RunFunc) *AppBuilder {
	b.app.runFunc = run
	return b
}

func (b *AppBuilder) WithPostRun(postRun RunFunc) *AppBuilder {
	b.app.postRunFunc = postRun
	return b
}

func (b *AppBuilder) WithNoVersion(noVersion bool) *AppBuilder {
	b.app.noVersion = noVersion
	return b
}

func (b *AppBuilder) WithNoConfig(noConfig bool) *AppBuilder {
	b.app.noConfig = noConfig
	return b
}

func (b *AppBuilder) WithOptions(options Options) *AppBuilder {
	b.app.options = options
	return b
}
