package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shim_alterCmd = &cobra.Command{
	Use:   "alter",
	Short: "alternate a shim's target source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shim_alterCmd).Standalone()
	shimCmd.AddCommand(shim_alterCmd)
}
