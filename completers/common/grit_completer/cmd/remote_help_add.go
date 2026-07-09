package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remote_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_help_addCmd).Standalone()

	remote_helpCmd.AddCommand(remote_help_addCmd)
}
