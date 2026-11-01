package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_localDisableCmd = &cobra.Command{
	Use:   "local-disable",
	Short: "Disable tailnet lock for this node only",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_localDisableCmd).Standalone()

	lockCmd.AddCommand(lock_localDisableCmd)
}
