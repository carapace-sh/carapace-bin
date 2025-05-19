package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Generates executable SQL from source, model, test, and analysis files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().Bool("defer", false, "Defer to the state variable for resolving unselected nodes")
	compileCmd.Flags().Bool("favor-state", false, "Defer to the state variable for resolving unselected nodes")
	compileCmd.Flags().BoolP("full-refresh", "f", false, "Dbt will drop incremental models and fully-recalculate")
	compileCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	compileCmd.Flags().String("log-path", "", "Configure the 'log-path'")
	compileCmd.Flags().Bool("no-defer", false, "Do not defer to the state variable for resolving unselected nodes")
	compileCmd.Flags().Bool("no-favor-state", false, "If defer is set, expect standard defer behaviour")
	compileCmd.Flags().Bool("no-version-check", false, "Skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	compileCmd.Flags().Bool("parse-only", false, "Only parse")
	compileCmd.Flags().String("selector", "", "The selector name to use")
	compileCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare")
	compileCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	compileCmd.Flags().String("target-path", "", "Configure the 'target-path'")
	compileCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	compileCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(compileCmd)
	addSelectionFlags(compileCmd)
	addModelsFlag(compileCmd)
	addProfileFlag(compileCmd)
	rootCmd.AddCommand(compileCmd)

	carapace.Gen(compileCmd).FlagCompletion(carapace.ActionMap{
		"log-path":    carapace.ActionFiles(),
		"state":       carapace.ActionDirectories(),
		"target-path": carapace.ActionDirectories(),
	})
}
