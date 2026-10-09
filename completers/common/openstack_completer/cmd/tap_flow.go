package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_flowCmd).Standalone()

	tapCmd.AddCommand(tap_flowCmd)
}
