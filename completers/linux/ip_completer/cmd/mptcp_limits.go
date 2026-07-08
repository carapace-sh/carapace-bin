package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mptcp_limitsCmd = &cobra.Command{
	Use:   "limits",
	Short: "get/change MPTCP subflow creation limits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mptcp_limitsCmd).Standalone()

	mptcpCmd.AddCommand(mptcp_limitsCmd)

	carapace.Gen(mptcp_limitsCmd).PositionalCompletion(
		carapace.ActionValues("set", "show"),
	)
}
