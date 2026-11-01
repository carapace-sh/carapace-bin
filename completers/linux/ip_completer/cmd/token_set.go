package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_setCmd = &cobra.Command{
	Use:   "set",
	Short: "set an interface token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_setCmd).Standalone()

	tokenCmd.AddCommand(token_setCmd)
}
