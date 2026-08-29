package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:     "agent",
	GroupID: "engine",
	Short:   "List available agents",
	Run:     func(cmd *cobra.Command, args []string) {},
	Aliases: []string{"agents"},
}

func init() {
	carapace.Gen(agentCmd).Standalone()

	agentCmd.Flags().String("output-format", "", "Output format (json, stream-json)")

	rootCmd.AddCommand(agentCmd)

	carapace.Gen(agentCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("json", "stream-json"),
	})
}
