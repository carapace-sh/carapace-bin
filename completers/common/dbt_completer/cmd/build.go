package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Run all Seeds, Models, Snapshots, and tests in DAG order",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("defer", false, "Defer to the state variable for resolving unselected nodes")
	buildCmd.Flags().BoolP("fail-fast", "x", false, "Stop execution upon a first failure")
	buildCmd.Flags().Bool("favor-state", false, "Defer to the state variable for resolving unselected nodes")
	buildCmd.Flags().BoolP("full-refresh", "f", false, "Drop incremental models and fully-recalculate")
	buildCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	buildCmd.Flags().Bool("indirect-selection", false, "Select all tests that are adjacent to selected resources")
	buildCmd.Flags().String("log-path", "", "Configure the 'log-path'")
	buildCmd.Flags().Bool("no-defer", false, "Do not defer to the state variable for resolving unselected nodes")
	buildCmd.Flags().Bool("no-favor-state", false, "Expect standard defer behaviour")
	buildCmd.Flags().Bool("no-version-check", false, "Skip ensuring dbt's version matches the one specified in the dbt_project.yml file")
	buildCmd.Flags().String("resource-type", "", "Limit by resource type")
	buildCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	buildCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare")
	buildCmd.Flags().Bool("store-failures", false, "Store test results (failing rows) in the database")
	buildCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	buildCmd.Flags().String("target-path", "", "Configure the 'target-path'")
	buildCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	addProjectDirFlag(buildCmd)
	addSelectionFlags(buildCmd)
	addProfileFlag(buildCmd)
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"indirect-selection": carapace.ActionValues("eager", "cautious", "buildable"),
		"log-path":           carapace.ActionFiles(),
		"resource-type":      carapace.ActionValues("model", "seed", "snapshot", "test", "all"),
		"state":              carapace.ActionDirectories(),
		"target-path":        carapace.ActionDirectories(),
	})
}
