package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugcheckstateCmd = &cobra.Command{
	Use:    "debugcheckstate",
	Short:  "validate the correctness of the current dirstate",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugcheckstateCmd).Standalone()

	rootCmd.AddCommand(debugcheckstateCmd)
}
