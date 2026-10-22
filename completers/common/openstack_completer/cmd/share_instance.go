package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_instanceCmd).Standalone()

	shareCmd.AddCommand(share_instanceCmd)
}
