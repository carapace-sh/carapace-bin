package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_replicaCmd = &cobra.Command{
	Use:   "replica",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_replicaCmd).Standalone()

	shareCmd.AddCommand(share_replicaCmd)
}
