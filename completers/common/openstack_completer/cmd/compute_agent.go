package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compute_agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_agentCmd).Standalone()

	computeCmd.AddCommand(compute_agentCmd)
}
