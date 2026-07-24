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

	exitNode_suggestCmd.Flags().Bool("force-probe", false, "force a probe of exit nodes before suggesting")
	exitNodeCmd.AddCommand(exitNode_suggestCmd)
}
