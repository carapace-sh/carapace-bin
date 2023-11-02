package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show commit description and changes in a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_showCmd).Standalone()

	helpCmd.AddCommand(help_showCmd)
}
