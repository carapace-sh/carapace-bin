package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var graphDependenciesCmd = &cobra.Command{
	Use:     "graph-dependencies",
	Short:   "Prints the terragrunt dependency graph to stdout",
	GroupID: "terragrunt",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(graphDependenciesCmd).Standalone()

	rootCmd.AddCommand(graphDependenciesCmd)
}
