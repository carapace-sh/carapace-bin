package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Create an on-demand snapshot with optional message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_snapshotCmd).Standalone()

	helpCmd.AddCommand(help_snapshotCmd)
}
