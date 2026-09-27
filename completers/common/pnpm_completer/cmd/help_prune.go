package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove extraneous packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pruneCmd).Standalone()

	helpCmd.AddCommand(help_pruneCmd)
}
