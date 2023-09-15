package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CommandOption defines optional parameters for initializing the command structure.
type CommandOption func(*Command)

// RunCommandFunc defines the application's command startup callback function.
type RunCommandFunc func(args []string) error

// Command is a sub command structure of a cli application.
type Command struct {
	usage    string
	desc     string
	options  Options
	commands []*Command
	runFunc  RunCommandFunc
}

// NewCommand creates a new sub command instance based on the given command name and other options.
func NewCommand(usage string, desc string, opts ...CommandOption) *Command {
	c := &Command{
		usage: usage,
		desc:  desc,
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

// AddCommand adds sub command to the current command.
func (c *Command) AddCommand(cmd *Command) {
	c.commands = append(c.commands, cmd)
}

// AddCommands adds multiple sub commands to the current command.
func (c *Command) AddCommands(cmds ...*Command) {
	c.commands = append(c.commands, cmds...)
}

func (c *Command) CobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   c.usage,
		Short: c.desc,
	}
	//samples.SetOutput(os.Stdout)
	cmd.Flags().SortFlags = false
	if len(c.commands) > 0 {
		for _, command := range c.commands {
			cmd.AddCommand(command.CobraCommand())
		}
	}
	if c.runFunc != nil {
		cmd.Run = func(cmd *cobra.Command, args []string) {
			if err := c.runFunc(args); err != nil {
				fmt.Printf("%v %v\n", "Error:", err)
				os.Exit(1)
			}
		}
	}
	if c.options != nil {
		for _, f := range c.options.Flags().FlagSets {
			cmd.Flags().AddFlagSet(f)
		}
		// c.options.AddFlags(samples.Flags())
	}
	//addHelpCommandFlag(c.usage, samples.Flags())
	return cmd
}

// AddCommand adds sub command to the application.
func (a *App) AddCommand(cmd *Command) {
	a.commands = append(a.commands, cmd)
}

// AddCommands adds multiple sub commands to the application.
func (a *App) AddCommands(cmds ...*Command) {
	a.commands = append(a.commands, cmds...)
}

// WithCommandRunFunc is used to set the application's command startup callback
// function option.
func WithCommandRunFunc(run RunCommandFunc) CommandOption {
	return func(c *Command) {
		c.runFunc = run
	}
}
