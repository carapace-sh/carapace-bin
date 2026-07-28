package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add one or more trusted signing keys to tailnet lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_addCmd).Standalone()

	lockCmd.AddCommand(lock_addCmd)
}
