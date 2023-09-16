package app

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dtweaveio/goframe/app/flag"
	"github.com/dtweaveio/goframe/app/version"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (app *App) Run() {
	app.buildCommand()
	if err := app.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func (app *App) buildCommand() {
	cmd := &cobra.Command{
		Use:   FormatBaseName(app.name),
		Short: app.shortName,
		Long:  app.description,
		// stop printing usage when the command errors
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          app.args,
	}
	cmd.Flags().SortFlags = true
	flag.InitFlags(cmd.Flags())

	if app.preRunFunc != nil {
		cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
			return app.preRunFunc(app.name)
		}
	}

	if app.runFunc != nil {
		cmd.RunE = app.runCommand
	}

	if app.postRunFunc != nil {
		cmd.PostRunE = func(cmd *cobra.Command, args []string) error {
			return app.postRunFunc(app.name)
		}
	}

	var namedFlagSets flag.NamedFlagSets
	if app.options != nil {
		namedFlagSets = app.options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	if !app.noConfig {
		addConfigFlag(app.name, namedFlagSets.FlagSet("global"))
	}

	if !app.noVersion {
		command := NewCommand("version", "Print version information", WithCommandRunFunc(func(args []string) error {
			println(version.Get().String())
			return nil
		}))
		app.commands = append(app.commands, command)
	}

	if len(app.commands) > 0 {
		for _, c := range app.commands {
			cmd.AddCommand(c.CobraCommand())
		}
	}

	// add new global flagset to samples FlagSet
	flag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name())
	cmd.Flags().AddFlagSet(namedFlagSets.FlagSet("global"))

	app.cmd = cmd
}

func (app *App) runCommand(cmd *cobra.Command, args []string) error {
	// marshall conf
	if !app.noConfig {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		if err := viper.Unmarshal(app.options); err != nil {
			return err
		}
	}
	if app.options != nil {
		if err := app.applyOptionRules(); err != nil {
			return err
		}
	}

	flag.PrintFlags(cmd.Flags())
	return app.runFunc(app.name)
}

func (app *App) applyOptionRules() error {
	if err := app.options.Complete(); err != nil {
		return err
	}

	if errs := app.options.Validate(); len(errs) != 0 {
		var builder strings.Builder
		for _, e := range errs {
			builder.WriteString(e.Error())
		}

		return errors.New(builder.String())
	}

	return nil
}
