package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_remote_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_remote_addCmd).Standalone()

	help_remoteCmd.AddCommand(help_remote_addCmd)
}
