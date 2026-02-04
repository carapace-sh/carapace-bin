package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skill_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if installed GitButler skills are up to date with the CLI version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_checkCmd).Standalone()

	skill_checkCmd.Flags().BoolP("global", "g", false, "Only check global installations (in home directory)")
	skill_checkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	skill_checkCmd.Flags().BoolP("local", "l", false, "Only check local installations (in current repository)")
	skill_checkCmd.Flags().BoolP("update", "u", false, "Automatically update any outdated skills found")
	skillCmd.AddCommand(skill_checkCmd)
}
