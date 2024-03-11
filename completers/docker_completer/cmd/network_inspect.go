package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var network_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] NETWORK [NETWORK...]",
	Short: "Display detailed information on one or more networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_inspectCmd).Standalone()

	network_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	network_inspectCmd.Flags().BoolP("verbose", "v", false, "Verbose output for diagnostics")
	networkCmd.AddCommand(network_inspectCmd)

	carapace.Gen(network_inspectCmd).PositionalAnyCompletion(docker.ActionNetworks())
}
