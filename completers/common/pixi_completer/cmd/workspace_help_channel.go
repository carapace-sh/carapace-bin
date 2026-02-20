package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Commands to manage workspace channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_channelCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_channelCmd)
}
