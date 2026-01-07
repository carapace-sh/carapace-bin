package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Retrieve metadata-related information about the current repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_infoCmd).Standalone()

	repo_infoCmd.Flags().String("format", "", "output format")
	repo_infoCmd.Flags().BoolS("z", "z", false, "alias for --format=nul")
	repoCmd.AddCommand(repo_infoCmd)

	repo_infoCmd.MarkFlagsMutuallyExclusive("format", "z")

	carapace.Gen(repo_infoCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValuesDescribed(
			"keyvalue", "use the = character as the delimiter",
			"nul", "use a newline character as the delimiter",
		),
	})

	carapace.Gen(repo_infoCmd).PositionalAnyCompletion(
		carapace.ActionValuesDescribed(
			"layout.bare", "true if this is a bare repository, otherwise false",
			"layout.shallow", "true if this is a shallow repository, otherwise false",
			"object.format", "The object format (hash algorithm) used in the repository",
			"references.format", "The reference storage format",
		).FilterArgs(),
	)
}
