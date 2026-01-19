package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_ctlCmd = &cobra.Command{
	Use:   "ctl",
	Short: "control build records",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_ctlCmd).Standalone()

	debug_ctlCmd.Flags().Bool("delete", false, "Delete build record")
	debug_ctlCmd.Flags().Bool("pin", false, "Pin build so it will not be garbage collected")
	debug_ctlCmd.Flags().Bool("unpin", false, "Unpin build so it will be garbage collected")
	debugCmd.AddCommand(debug_ctlCmd)
}
