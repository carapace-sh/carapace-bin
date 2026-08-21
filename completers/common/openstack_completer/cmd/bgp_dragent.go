package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_dragentCmd = &cobra.Command{
	Use:   "dragent",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_dragentCmd).Standalone()

	bgpCmd.AddCommand(bgp_dragentCmd)
}
