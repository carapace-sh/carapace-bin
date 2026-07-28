package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_breakTcpConnsCmd = &cobra.Command{
	Use:   "break-tcp-conns",
	Short: "Break any open TCP connections from the daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_breakTcpConnsCmd).Standalone()

	debugCmd.AddCommand(debug_breakTcpConnsCmd)
}
