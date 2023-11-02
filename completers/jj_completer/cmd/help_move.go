package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move changes from one revision into another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_moveCmd).Standalone()

	helpCmd.AddCommand(help_moveCmd)
}
