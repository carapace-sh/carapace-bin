package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_resize_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert server resize.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_resize_revertCmd).Standalone()

	server_resizeCmd.AddCommand(server_resize_revertCmd)
}
