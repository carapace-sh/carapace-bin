package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_instance_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_instance_exportCmd).Standalone()

	share_snapshot_instanceCmd.AddCommand(share_snapshot_instance_exportCmd)
}
