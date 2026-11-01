package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_quotaCmd = &cobra.Command{
	Use:   "quota",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_quotaCmd).Standalone()

	shareCmd.AddCommand(share_quotaCmd)
}
