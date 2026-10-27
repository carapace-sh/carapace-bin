package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remote_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a remote and its remote-tracking refs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_help_removeCmd).Standalone()

	remote_helpCmd.AddCommand(remote_help_removeCmd)
}
