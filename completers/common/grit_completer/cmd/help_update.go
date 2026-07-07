package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update grit to the latest release (re-runs the install script)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_updateCmd).Standalone()

	helpCmd.AddCommand(help_updateCmd)
}
