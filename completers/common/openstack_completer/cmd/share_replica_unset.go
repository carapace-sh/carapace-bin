package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_replica_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a share replica property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replica_unsetCmd).Standalone()

	share_replica_unsetCmd.Flags().String("property", "", "Remove a property from replica (repeat option to remove multiple properties)")
	share_replicaCmd.AddCommand(share_replica_unsetCmd)
}
