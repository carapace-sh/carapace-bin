package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restores this share or more shares from the recycle bin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_restoreCmd).Standalone()

	shareCmd.AddCommand(share_restoreCmd)
}
