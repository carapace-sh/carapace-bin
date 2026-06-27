package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exitNode_suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest the best available exit node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitNode_suggestCmd).Standalone()

	exitNodeCmd.AddCommand(exitNode_suggestCmd)
}
