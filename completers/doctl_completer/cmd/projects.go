package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manage projects and assign resources to them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectsCmd).Standalone()
	rootCmd.AddCommand(projectsCmd)
}
