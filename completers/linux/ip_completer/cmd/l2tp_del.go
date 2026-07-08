package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var l2tp_delCmd = &cobra.Command{
	Use:     "del",
	Aliases: []string{"delete"},
	Short:   "destroy an L2TP tunnel or session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(l2tp_delCmd).Standalone()

	l2tpCmd.AddCommand(l2tp_delCmd)

	carapace.Gen(l2tp_delCmd).PositionalCompletion(
		carapace.ActionValues("tunnel", "session"),
	)
}
