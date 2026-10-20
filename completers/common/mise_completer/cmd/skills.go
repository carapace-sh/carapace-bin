package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillsCmd = &cobra.Command{
	Use:     "skills",
	Short:   "List and link the agent skills that active tools ship",
	Aliases: []string{"skill"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillsCmd).Standalone()

	skillsCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(skillsCmd)
}
