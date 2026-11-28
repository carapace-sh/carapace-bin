package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateDependenciesCmd = &cobra.Command{
	Use:   "update-dependencies",
	Short: "Update dependencies of the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateDependenciesCmd).Standalone()
	rootCmd.AddCommand(updateDependenciesCmd)
}
