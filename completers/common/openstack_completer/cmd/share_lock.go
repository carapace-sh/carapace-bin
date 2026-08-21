package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_lockCmd).Standalone()

	shareCmd.AddCommand(share_lockCmd)
}
