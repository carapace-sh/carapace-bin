package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var suspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "suspends the machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspendCmd).Standalone()

	suspendCmd.Flags().BoolP("all-global", "a", false, "Suspend all running vms globally.")
	rootCmd.AddCommand(suspendCmd)

	carapace.Gen(suspendCmd).PositionalCompletion(
		vagrant.ActionMachines(),
	)
}
