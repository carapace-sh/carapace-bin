package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_channel_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a channel to the manifest and updates the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_channel_help_addCmd).Standalone()

	workspace_channel_helpCmd.AddCommand(workspace_channel_help_addCmd)
}
