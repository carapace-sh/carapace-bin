package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xfrm_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "manage Security Policy Database (SPD)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xfrm_policyCmd).Standalone()

	xfrmCmd.AddCommand(xfrm_policyCmd)

	carapace.Gen(xfrm_policyCmd).PositionalCompletion(
		carapace.ActionValues("add", "update", "delete", "get", "deleteall", "list", "flush", "count", "set", "setdefault", "getdefault"),
	)
}
