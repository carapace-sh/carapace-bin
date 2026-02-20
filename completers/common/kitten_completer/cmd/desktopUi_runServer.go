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
