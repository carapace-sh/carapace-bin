package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert a share to the specified snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_revertCmd).Standalone()

	share_revertCmd.Flags().Bool("wait", false, "Wait for share revert")
	shareCmd.AddCommand(share_revertCmd)
}
