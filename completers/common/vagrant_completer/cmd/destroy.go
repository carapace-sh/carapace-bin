package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "stops and deletes all traces of the vagrant machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()

	destroyCmd.Flags().BoolP("force", "f", false, "Destroy without confirmation.")
	destroyCmd.Flags().BoolP("graceful", "g", false, "Gracefully poweroff of VM")
	destroyCmd.Flags().Bool("no-parallel", false, "Disable parallelism if provider supports it (automatically enables force)")
	destroyCmd.Flags().Bool("parallel", false, "Enable parallelism if provider supports it (automatically enables force)")
	rootCmd.AddCommand(destroyCmd)

	carapace.Gen(destroyCmd).PositionalCompletion(
		vagrant.ActionMachines(),
	)
}
