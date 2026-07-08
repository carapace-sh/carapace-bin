package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"show"},
	Short:   "list all interface tokens",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_listCmd).Standalone()

	tokenCmd.AddCommand(token_listCmd)
}
