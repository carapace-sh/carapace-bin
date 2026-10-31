package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keys_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Return the value of a stored key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keys_getCmd).Standalone()

	keysCmd.AddCommand(keys_getCmd)
}
