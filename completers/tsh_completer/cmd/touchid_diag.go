package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var touchid_diagCmd = &cobra.Command{
	Use:    "diag",
	Short:  "Run Touch ID diagnostics",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(touchid_diagCmd).Standalone()

	touchidCmd.AddCommand(touchid_diagCmd)
}
