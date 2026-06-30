package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var funnel_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset current funnel config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(funnel_resetCmd).Standalone()

	funnelCmd.AddCommand(funnel_resetCmd)
}
