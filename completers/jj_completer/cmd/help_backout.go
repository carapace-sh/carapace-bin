package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_backoutCmd = &cobra.Command{
	Use:   "backout",
	Short: "Apply the reverse of a revision on top of another revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_backoutCmd).Standalone()

	helpCmd.AddCommand(help_backoutCmd)
}
