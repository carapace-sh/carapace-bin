package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_serviceCmd).Standalone()

	shareCmd.AddCommand(share_serviceCmd)
}
