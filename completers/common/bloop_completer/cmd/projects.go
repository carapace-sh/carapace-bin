package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "list projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectsCmd).Standalone()

	projectsCmd.Flags().Bool("dot-graph", false, "print out a dot graph you can pipe into `dot`")
	rootCmd.AddCommand(projectsCmd)
}
