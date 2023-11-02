package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_filesCmd = &cobra.Command{
	Use:   "files",
	Short: "List files in a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_filesCmd).Standalone()

	helpCmd.AddCommand(help_filesCmd)
}
