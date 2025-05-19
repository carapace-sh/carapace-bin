package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Show some helpful information about dbt for debugging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()

	debugCmd.Flags().Bool("config-dir", false, "DBT will show path information for this project")
	debugCmd.Flags().Bool("no-version-check", false, "skip ensuring dbt's version matches the one specified in the dbt_project.yml file")
	debugCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	debugCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(debugCmd)
	addProfileFlag(debugCmd)
	rootCmd.AddCommand(debugCmd)
}
