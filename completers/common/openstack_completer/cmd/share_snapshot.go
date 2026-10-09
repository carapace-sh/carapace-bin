package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshotCmd).Standalone()

	shareCmd.AddCommand(share_snapshotCmd)
}
