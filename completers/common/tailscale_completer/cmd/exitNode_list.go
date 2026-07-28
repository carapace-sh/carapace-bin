package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exitNode_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show exit nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitNode_listCmd).Standalone()

	exitNode_listCmd.Flags().String("filter", "", "filter exit nodes by country")
	exitNodeCmd.AddCommand(exitNode_listCmd)
}
