package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_abandonCmd = &cobra.Command{
	Use:   "abandon",
	Short: "Abandon a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_abandonCmd).Standalone()

	helpCmd.AddCommand(help_abandonCmd)
}
