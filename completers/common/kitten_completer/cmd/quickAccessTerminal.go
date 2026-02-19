package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quickAccessTerminalCmd = &cobra.Command{
	Use:   "quick-access-terminal",
	Short: "A quick access terminal window that you can bring up instantly with a keypress or a command.",
}

func init() {
	rootCmd.AddCommand(quickAccessTerminalCmd)
	carapace.Gen(quickAccessTerminalCmd).Standalone()

	quickAccessTerminalCmd.Flags().StringArrayP("config", "c", nil, "Specify a path to the configuration file(s) to use. All configuration files are merged onto the builtin quick-access-terminal.conf, overriding the builtin values. This option can be specified multiple times to read multiple configuration files in sequence, which are merged. Use the special value NONE to not load any config file.")
	quickAccessTerminalCmd.Flags().StringArrayP("override", "o", nil, "Override individual configuration options, can be specified multiple times. Syntax: name=value. For example: -o lines=12")
	quickAccessTerminalCmd.Flags().Bool("detach", false, "Detach from the controlling terminal, if any, running in an independent child process, the parent process exits immediately.")
	quickAccessTerminalCmd.Flags().String("detached-log", "", "Path to a log file to store STDOUT/STDERR when using --detach")
	quickAccessTerminalCmd.Flags().String("instance-group", "", "The unique name of this quick access terminal Use a different name if you want multiple such terminals.")
	quickAccessTerminalCmd.Flags().Bool("debug-rendering", false, "For debugging interactions with the compositor/window manager.")
	quickAccessTerminalCmd.Flags().Bool("debug-input", false, "For debugging interactions with the compositor/window manager.")
	quickAccessTerminalCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(quickAccessTerminalCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.Batch(carapace.ActionFiles(), carapace.ActionValues("NONE", "-", "/dev/stdin")).ToA(),
	})

	carapace.Gen(quickAccessTerminalCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
