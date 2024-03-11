package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var list_expiredCmd = &cobra.Command{
	Use:   "expired",
	Short: "List all posts already expired",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_expiredCmd).Standalone()
	listCmd.AddCommand(list_expiredCmd)
}
