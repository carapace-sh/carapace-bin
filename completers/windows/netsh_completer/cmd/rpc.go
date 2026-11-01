package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "RPC configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rpcCmd).Standalone()
	rootCmd.AddCommand(rpcCmd)
}
