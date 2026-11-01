package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillsCmd = &cobra.Command{
	Use:   "skills <command>",
	Short: "Manage glab agent skills. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillsCmd).Standalone()

	rootCmd.AddCommand(skillsCmd)
}
