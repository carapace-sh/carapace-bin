package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var l2tp_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "show L2TP tunnel or session information",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(l2tp_showCmd).Standalone()

	l2tpCmd.AddCommand(l2tp_showCmd)

	carapace.Gen(l2tp_showCmd).PositionalCompletion(
		carapace.ActionValues("tunnel", "session"),
	)
}
