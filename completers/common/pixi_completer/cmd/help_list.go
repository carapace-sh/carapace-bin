package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the packages of the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_listCmd).Standalone()

	helpCmd.AddCommand(help_listCmd)
}
