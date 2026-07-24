package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_break_tcp_connsCmd = &cobra.Command{
	Use:   "break-tcp-conns",
	Short: "Break any open TCP connections from the daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_break_tcp_connsCmd).Standalone()

	debugCmd.AddCommand(debug_break_tcp_connsCmd)
}
