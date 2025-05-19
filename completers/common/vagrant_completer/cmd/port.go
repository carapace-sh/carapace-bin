package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "displays information about guest port mappings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(portCmd).Standalone()

	portCmd.Flags().String("guest", "", "Output the host port that maps to the given guest port")
	rootCmd.AddCommand(portCmd)

	// TODO guest port completion

	carapace.Gen(portCmd).PositionalCompletion(
		vagrant.ActionMachines(),
	)
}
