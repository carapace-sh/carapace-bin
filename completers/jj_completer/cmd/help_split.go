package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split a revision in two",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_splitCmd).Standalone()

	helpCmd.AddCommand(help_splitCmd)
}
