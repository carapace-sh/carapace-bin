package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_link_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the link at the given point in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_link_removeCmd).Standalone()

	help_linkCmd.AddCommand(help_link_removeCmd)
}
