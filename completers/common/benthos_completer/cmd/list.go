package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Benthos component types",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("format", "", "Print the component list in a specific format. Options are text, json or cue.")
	listCmd.Flags().String("status", "", "Filter the component list to only those matching the given status. Options are stable, beta or experimental.")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json", "cue"),
		"status": carapace.ActionValues("stable", "beta", "experimental"),
	})

	carapace.Gen(listCmd).PositionalAnyCompletion(
		carapace.ActionValues(
			"bloblang-functions",
			"bloblang-methods",
			"buffers",
			"caches",
			"inputs",
			"metrics",
			"outputs",
			"processors",
			"rate-limits",
			"tracers",
		).FilterArgs(),
	)
}
