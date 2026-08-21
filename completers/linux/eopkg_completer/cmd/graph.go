package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "produce a dependency graph for the given package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(graphCmd).Standalone()

	graphCmd.Flags().Bool("ignore-installed", false, "ignore currently installed packages in the graph")
	graphCmd.Flags().Bool("installed", false, "only include installed packages in the graph")
	graphCmd.Flags().StringP("output", "o", "", "override path to the output file")
	graphCmd.Flags().StringP("repository", "r", "", "only consider packages from the given repository")

	rootCmd.AddCommand(graphCmd)
}
