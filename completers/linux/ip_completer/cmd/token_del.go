package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete an interface token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_delCmd).Standalone()

	tokenCmd.AddCommand(token_delCmd)
}
