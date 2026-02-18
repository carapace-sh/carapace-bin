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
	inspectCmd.Flags().BoolS("d", "d", false, "Specify the full path to the subdirectory if you don't want to inspect the whole project")
	inspectCmd.Flags().StringS("format", "format", "", "Specify the format of the output file with inspection results")
	inspectCmd.Flags().BoolS("v", "v", false, "Set the verbosity level of the output")

	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).FlagCompletion(carapace.ActionMap{
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
