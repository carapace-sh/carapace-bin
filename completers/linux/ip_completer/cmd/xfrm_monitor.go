package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xfrm_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "state monitoring for xfrm objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xfrm_monitorCmd).Standalone()

	xfrmCmd.AddCommand(xfrm_monitorCmd)
}
