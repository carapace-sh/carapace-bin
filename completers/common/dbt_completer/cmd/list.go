package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the resources in your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	listCmd.Flags().String("indirect-selection", "", "Select all tests that are adjacent to selected resources")
	listCmd.Flags().String("output", "", "Output format")
	listCmd.Flags().String("output-keys", "", "Output keys")
	listCmd.Flags().String("resource-type", "", "Limit to resource type")
	listCmd.Flags().String("selector", "", "The selector name to use")
	listCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare")
	listCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	listCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(listCmd)
	addSelectionFlags(listCmd)
	addModelsFlag(listCmd)
	addProfileFlag(listCmd)
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"output":        carapace.ActionValues("json", "name", "path", "selector"),
		"resource-type": carapace.ActionValues("analysis", "snapshot", "source", "metric", "model", "seed", "test", "exposure", "default", "all"),
		"state":         carapace.ActionDirectories(),
	})
}
