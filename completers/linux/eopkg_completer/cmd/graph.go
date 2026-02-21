package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Graph package relations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(graphCmd).Standalone()

	graphCmd.Flags().BoolP("repository", "R", false, "Graph only from repository")
	graphCmd.Flags().Bool("installed", false, "Graph only installed packages")
	graphCmd.Flags().Bool("ignore-installed", false, "Do not show installed packages")
	graphCmd.Flags().StringP("output", "o", "", "Output file for graph")

	rootCmd.AddCommand(graphCmd)

	carapace.Gen(graphCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
