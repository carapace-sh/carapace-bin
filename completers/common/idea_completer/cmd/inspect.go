package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Perform code inspection on the specified project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	inspectCmd.Flags().BoolS("changes", "changes", false, "Run inspections only on local uncommitted changes")
	inspectCmd.Flags().StringS("d", "d", "", "Specify the full path to the subdirectory if you don't want to inspect the whole project")
	inspectCmd.Flags().StringS("format", "format", "xml", "Specify the format of the output file with inspection results")
	inspectCmd.Flags().BoolS("v0", "v0", true, "Set the verbosity level of the output")
	inspectCmd.Flags().BoolS("v1", "v1", false, "Set the verbosity level of the output")
	inspectCmd.Flags().BoolS("v2", "v2", false, "Set the verbosity level of the output")

	inspectCmd.MarkFlagsMutuallyExclusive("v0", "v1", "v2")

	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).FlagCompletion(carapace.ActionMap{
		"d": carapace.ActionDirectories(),
		"format": carapace.ActionValuesDescribed(
			"xml", "default",
			"json", "",
			"plain", "",
		),
	})

	carapace.Gen(inspectCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
