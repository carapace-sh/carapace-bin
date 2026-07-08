package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var l2tp_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new L2TP tunnel or session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(l2tp_addCmd).Standalone()

	l2tpCmd.AddCommand(l2tp_addCmd)

	carapace.Gen(l2tp_addCmd).PositionalCompletion(
		carapace.ActionValues("tunnel", "session"),
	)
}
