package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_link_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all links in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_link_listCmd).Standalone()

	help_linkCmd.AddCommand(help_link_listCmd)
}
