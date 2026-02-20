package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runServerCmd = &cobra.Command{
	Use:   "run-server",
	Short: "Start the various servers used to integrate with the Linux desktop",
}

func init() {
	runServerCmd.AddCommand(desktopUiCmd)
	carapace.Gen(runServerCmd).Standalone()

	runServerCmd.Flags().StringP("config", "c", "", "Specify a path to the configuration file(s) to use")
	runServerCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	runServerCmd.Flags().StringP("override", "o", "", "Override individual configuration options, can be specified multiple times")

	carapace.Gen(runServerCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}

// Usage: kitten desktop-ui run-server
//
// This should be run very early in the startup sequence of your window manager,
// before any other programs are run.
//
// Options:
//   --override, -o
//     Override individual configuration options, can be specified multiple times.
//     Syntax: name=value. For example: -o color_scheme=dark
//
//   --config, -c
//     Specify a path to the configuration file(s) to use. All configuration files
//     are merged onto the builtin desktop-ui-portal.conf, overriding the builtin
//     values. This option can be specified multiple times to read multiple
//     configuration files in sequence, which are merged. Use the special value
//     NONE to not load any config file.
//
//   --help, -h [=no]
//     Show help for this command
