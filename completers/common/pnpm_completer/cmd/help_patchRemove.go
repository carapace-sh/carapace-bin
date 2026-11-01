package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_patchRemoveCmd = &cobra.Command{
	Use:   "patch-remove",
	Short: "Remove existing patch files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_patchRemoveCmd).Standalone()

	helpCmd.AddCommand(help_patchRemoveCmd)
}
