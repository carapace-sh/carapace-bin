package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete all folders in the clean-targets list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	cleanCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(cleanCmd)
	addProfileFlag(cleanCmd)
	rootCmd.AddCommand(cleanCmd)
}
