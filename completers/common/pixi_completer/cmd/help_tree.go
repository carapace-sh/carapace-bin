package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show a tree of workspace dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_treeCmd).Standalone()

	helpCmd.AddCommand(help_treeCmd)
}
