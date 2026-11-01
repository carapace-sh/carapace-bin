package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_abandonCmd = &cobra.Command{
	Use:   "abandon",
	Short: "Abandon a share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_abandonCmd).Standalone()

	share_abandonCmd.Flags().Bool("wait", false, "Wait until share is abandoned")
	shareCmd.AddCommand(share_abandonCmd)
}
