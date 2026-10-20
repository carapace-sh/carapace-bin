package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skills_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List the skills the active tools declare",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skills_lsCmd).Standalone()

	skills_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	skills_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	skillsCmd.AddCommand(skills_lsCmd)
}
