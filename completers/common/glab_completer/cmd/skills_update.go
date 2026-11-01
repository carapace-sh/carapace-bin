package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skills_updateCmd = &cobra.Command{
	Use:   "update [name]",
	Short: "Update installed agent skills to the current shipped version. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skills_updateCmd).Standalone()

	skills_updateCmd.Flags().Bool("all", false, "Update every installed skill.")
	skillsCmd.AddCommand(skills_updateCmd)
}
