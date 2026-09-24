package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_qosCmd = &cobra.Command{
	Use:   "qos",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_qosCmd).Standalone()

	shareCmd.AddCommand(share_qosCmd)
}
