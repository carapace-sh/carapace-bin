package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopUi_runServerCmd = &cobra.Command{
	Use:   "run-server",
	Short: "Start the various servers used to integrate with the Linux desktop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desktopUi_runServerCmd).Standalone()

	desktopUi_runServerCmd.Flags().StringP("config", "c", "", "Specify a path to the configuration file(s) to use")
	desktopUi_runServerCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	desktopUi_runServerCmd.Flags().StringP("override", "o", "", "Override individual configuration options, can be specified multiple times")
	desktopUiCmd.AddCommand(desktopUi_runServerCmd)

	carapace.Gen(desktopUi_runServerCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
