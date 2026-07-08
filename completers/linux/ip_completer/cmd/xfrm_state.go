package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xfrm_stateCmd = &cobra.Command{
	Use:   "state",
	Short: "manage Security Association Database (SAD)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xfrm_stateCmd).Standalone()

	xfrmCmd.AddCommand(xfrm_stateCmd)

	carapace.Gen(xfrm_stateCmd).PositionalCompletion(
		carapace.ActionValues("add", "update", "allocspi", "delete", "get", "deleteall", "list", "flush", "count"),
	)
}
