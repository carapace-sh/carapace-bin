package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new DBT project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("skip-profile-setup", "s", false, "Skip interative profile setup.")
	initCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	initCmd.Flags().String("vars", "", "Supply variables to the project.")
	addProjectDirFlag(initCmd)
	addProfileFlag(initCmd)
	rootCmd.AddCommand(initCmd)
}
