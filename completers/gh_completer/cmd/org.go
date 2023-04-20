package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orgCmd = &cobra.Command{
	Use:     "org <command>",
	Short:   "Manage organizations",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orgCmd).Standalone()
	orgCmd.AddGroup(
		&cobra.Group{ID: "General commands", Title: "General commands"},
	)

	rootCmd.AddCommand(orgCmd)
}
