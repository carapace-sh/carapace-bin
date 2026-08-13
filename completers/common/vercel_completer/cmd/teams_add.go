package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var teams_addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create"},
	Short:   "Create a new team",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_addCmd).Standalone()

	teams_addCmd.Flags().String("name", "", "Name of the team")
	teams_addCmd.Flags().String("slug", "", "Slug for the team")

	teamsCmd.AddCommand(teams_addCmd)
}
