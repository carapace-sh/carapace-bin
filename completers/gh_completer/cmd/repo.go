package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:     "repo <command>",
	Short:   "Manage repositories",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()
	repoCmd.AddGroup(
		&cobra.Group{ID: "general", Title: "General commands"},
		&cobra.Group{ID: "targeted", Title: "Targeted commands"},
	)

	rootCmd.AddCommand(repoCmd)
}
