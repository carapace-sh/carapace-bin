package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var link_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all links in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_listCmd).Standalone()

	link_listCmd.Flags().BoolP("help", "h", false, "Print help")
	link_listCmd.Flags().Bool("staged", false, "Only show links with staged changes")
	linkCmd.AddCommand(link_listCmd)
}
