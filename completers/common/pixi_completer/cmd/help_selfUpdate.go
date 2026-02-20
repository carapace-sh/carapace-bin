package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update pixi to the latest version or a specific version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_selfUpdateCmd).Standalone()

	helpCmd.AddCommand(help_selfUpdateCmd)
}
