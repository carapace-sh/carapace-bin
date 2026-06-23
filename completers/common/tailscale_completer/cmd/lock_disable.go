package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Consume a disablement secret to shut down tailnet lock for the tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_disableCmd).Standalone()

	lockCmd.AddCommand(lock_disableCmd)
}
