package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_derpCmd = &cobra.Command{
	Use:   "derp",
	Short: "Test a DERP configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_derpCmd).Standalone()

	debugCmd.AddCommand(debug_derpCmd)
}
