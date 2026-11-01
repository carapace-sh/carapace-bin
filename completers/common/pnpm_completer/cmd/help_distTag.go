package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_distTagCmd = &cobra.Command{
	Use:   "dist-tag",
	Short: "Manage a package's distribution tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_distTagCmd).Standalone()

	helpCmd.AddCommand(help_distTagCmd)
}
