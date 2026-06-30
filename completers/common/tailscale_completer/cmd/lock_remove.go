package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove one or more trusted signing keys from tailnet lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_removeCmd).Standalone()

	lockCmd.AddCommand(lock_removeCmd)
}
