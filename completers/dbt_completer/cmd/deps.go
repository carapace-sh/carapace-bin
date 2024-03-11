package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Pull the most recent version of the dependencies listed in packages.yml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(depsCmd).Standalone()

	depsCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	depsCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(depsCmd)
	addProfileFlag(depsCmd)
	rootCmd.AddCommand(depsCmd)
}
