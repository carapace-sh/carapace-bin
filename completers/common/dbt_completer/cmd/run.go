package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Compile SQL and execute against the current target database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("defer", false, "Defer to the state variable for resolving unselected nodes")
	runCmd.Flags().BoolP("fail-fast", "x", false, "Stop execution upon a first failure")
	runCmd.Flags().Bool("favor-state", false, "Defer to the state variable for resolving unselected nodes")
	runCmd.Flags().BoolP("full-refresh", "f", false, "Drop incremental models and fully-recalculate")
	runCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	runCmd.Flags().String("log-path", "", "Configure the 'log-path'")
	runCmd.Flags().Bool("no-defer", false, "Do not defer to the state variable for resolving unselected nodes")
	runCmd.Flags().Bool("no-favor-state", false, "If defer is set, expect standard defer behaviour")
	runCmd.Flags().Bool("no-version-check", false, "Skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	runCmd.Flags().String("selector", "", "The selector name to use")
	runCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare")
	runCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	runCmd.Flags().String("target-path", "", "Configure the 'target-path'")
	runCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	runCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(runCmd)
	addSelectionFlags(runCmd)
	addModelsFlag(runCmd)
	addProfileFlag(runCmd)
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"log-path":    carapace.ActionFiles(),
		"state":       carapace.ActionDirectories(),
		"target-path": carapace.ActionDirectories(),
	})
}
