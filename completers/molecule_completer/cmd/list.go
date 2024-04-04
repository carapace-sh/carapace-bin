package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/molecule_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List status of instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd)

	listCmd.Flags().StringP("format", "f", "simple", "Change output format")
	listCmd.Flags().Bool("help", false, "Show help message and exit")
	listCmd.Flags().StringP("scenario-name", "s", "default", "Name of the scenario to target")

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format":        carapace.ActionValues("plain", "simple", "yaml"),
		"scenario-name": action.ActionScenarios(),
	})

	rootCmd.AddCommand(listCmd)
}
