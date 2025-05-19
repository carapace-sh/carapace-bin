package cmd

import (
	"github.com/carapace-sh/carapace"
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
		&cobra.Group{ID: "Targeted commands", Title: ""},
		&cobra.Group{ID: "General commands", Title: ""},
	)

	rootCmd.AddCommand(repoCmd)
}
