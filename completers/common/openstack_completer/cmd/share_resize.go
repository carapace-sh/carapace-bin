package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize a share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_resizeCmd).Standalone()

	share_resizeCmd.Flags().Bool("force", false, "Only applicable when increasing the size of the share，only available with microversion 2.64 and higher.")
	share_resizeCmd.Flags().Bool("wait", false, "Wait for share resize")
	shareCmd.AddCommand(share_resizeCmd)
}
