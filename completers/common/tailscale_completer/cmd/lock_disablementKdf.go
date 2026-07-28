package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_disablementKdfCmd = &cobra.Command{
	Use:   "disablement-kdf",
	Short: "Compute a disablement value from a disablement secret (advanced users only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_disablementKdfCmd).Standalone()

	lockCmd.AddCommand(lock_disablementKdfCmd)
}
