package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var list_draftsCmd = &cobra.Command{
	Use:   "drafts",
	Short: "List all drafts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_draftsCmd).Standalone()
	listCmd.AddCommand(list_draftsCmd)
}
