package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List project directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectsCmd).Standalone()

	projectsCmd.Flags().Bool("json", false, "Output as JSON")
	rootCmd.AddCommand(projectsCmd)
}
