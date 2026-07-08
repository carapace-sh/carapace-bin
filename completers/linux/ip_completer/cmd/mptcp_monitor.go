package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mptcp_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "display MPTCP connection events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mptcp_monitorCmd).Standalone()

	mptcpCmd.AddCommand(mptcp_monitorCmd)
}
