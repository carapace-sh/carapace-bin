package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var link_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all links in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_help_listCmd).Standalone()

	link_helpCmd.AddCommand(link_help_listCmd)
}
