package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var piv_agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start PIV key agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(piv_agentCmd).Standalone()

	pivCmd.AddCommand(piv_agentCmd)
}
