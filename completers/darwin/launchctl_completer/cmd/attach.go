package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach the system's debugger to a service",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()
	rootCmd.AddCommand(attachCmd)
	carapace.Gen(attachCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
