package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_obslogCmd = &cobra.Command{
	Use:   "obslog",
	Short: "Show how a change has evolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_obslogCmd).Standalone()

	helpCmd.AddCommand(help_obslogCmd)
}
