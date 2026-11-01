package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Generate an AGENTS.md file with Vercel deployment best practices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentCmd).Standalone()

	agentCmd.Flags().Bool("yes", false, "Skip confirmation")

	rootCmd.AddCommand(agentCmd)

	carapace.Gen(agentCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
