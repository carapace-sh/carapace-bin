package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var list_allCmd = &cobra.Command{
	Use:   "all",
	Short: "List all posts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_allCmd).Standalone()
	listCmd.AddCommand(list_allCmd)
}
