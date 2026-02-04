package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "Manage Claude AI skills for GitButler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillCmd).Standalone()

	skillCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(skillCmd)
}
