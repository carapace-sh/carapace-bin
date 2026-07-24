package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a node or pre-approved auth key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_signCmd).Standalone()

	lockCmd.AddCommand(lock_signCmd)
}
