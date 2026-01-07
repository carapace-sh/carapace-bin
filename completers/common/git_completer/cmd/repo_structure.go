package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_structureCmd = &cobra.Command{
	Use:   "structure",
	Short: "Retrieve statistics about the current repository structure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_structureCmd).Standalone()

	repo_structureCmd.Flags().String("format", "", "output format")
	repoCmd.AddCommand(repo_structureCmd)

	carapace.Gen(repo_structureCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValuesDescribed(
			"table", "Outputs repository stats in a human-friendly table",
			"keyvalue", "Each line of output contains a key-value pair for a repository stat",
			"nul", "Similar to keyvalue, but uses a NUL character to delimit between key-value pairs",
		),
	})
}
