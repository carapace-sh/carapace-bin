package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_replica_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replica_exportCmd).Standalone()

	share_replicaCmd.AddCommand(share_replica_exportCmd)
}
