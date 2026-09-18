package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_transferCmd).Standalone()

	shareCmd.AddCommand(share_transferCmd)
}
