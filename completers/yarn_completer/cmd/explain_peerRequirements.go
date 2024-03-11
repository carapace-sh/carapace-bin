package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var explain_peerRequirementsCmd = &cobra.Command{
	Use:   "peer-requirements",
	Short: "Explain a set of peer requirements",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explain_peerRequirementsCmd).Standalone()

	explainCmd.AddCommand(explain_peerRequirementsCmd)
}
