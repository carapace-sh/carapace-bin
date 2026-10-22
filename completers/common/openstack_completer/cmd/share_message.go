package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_messageCmd = &cobra.Command{
	Use:   "message",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_messageCmd).Standalone()

	shareCmd.AddCommand(share_messageCmd)
}
