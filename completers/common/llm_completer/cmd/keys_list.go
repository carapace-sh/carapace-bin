package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keys_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List names of all stored keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keys_listCmd).Standalone()

	keysCmd.AddCommand(keys_listCmd)
}
