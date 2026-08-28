package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentlogCmd = &cobra.Command{
	Use:    "agentlog",
	Short:  "AI: capture agent logs into GitMeta",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentlogCmd).Standalone()

	agentlogCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(agentlogCmd)
}
