package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var defaultAgentCmd = &cobra.Command{
	Use:   "default-agent",
	Short: "Set agent as the default one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(defaultAgentCmd).Standalone()
	rootCmd.AddCommand(defaultAgentCmd)
}
