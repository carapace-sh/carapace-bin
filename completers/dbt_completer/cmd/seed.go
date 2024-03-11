package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Load data from csv files into your data warehouse",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(seedCmd).Standalone()

	seedCmd.Flags().BoolP("full-refresh", "f", false, "Drop existing seed tables and recreate them")
	seedCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	seedCmd.Flags().String("log-path", "", "Configure the 'log-path'")
	seedCmd.Flags().Bool("no-version-check", false, "Skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	seedCmd.Flags().String("selector", "", "The selector name to use")
	seedCmd.Flags().Bool("show", false, "Show a sample of the loaded data in the terminal")
	seedCmd.Flags().String("state", "", "Jse the given directory as the source for json files to compare")
	seedCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	seedCmd.Flags().String("target-path", "", "Configure the 'target-path'")
	seedCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	seedCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(seedCmd)
	addSelectionFlags(seedCmd)
	addModelsFlag(seedCmd)
	addProfileFlag(seedCmd)
	rootCmd.AddCommand(seedCmd)

	carapace.Gen(seedCmd).FlagCompletion(carapace.ActionMap{
		"log-path":    carapace.ActionFiles(),
		"state":       carapace.ActionDirectories(),
		"target-path": carapace.ActionDirectories(),
	})
}
