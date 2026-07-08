package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_getCmd = &cobra.Command{
	Use:   "get",
	Short: "get the interface token for a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_getCmd).Standalone()

	tokenCmd.AddCommand(token_getCmd)
}
