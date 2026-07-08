package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mptcp_endpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "manage MPTCP endpoints",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mptcp_endpointCmd).Standalone()

	mptcpCmd.AddCommand(mptcp_endpointCmd)

	carapace.Gen(mptcp_endpointCmd).PositionalCompletion(
		carapace.ActionValues("add", "delete", "change", "show", "flush"),
	)
}
