package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_inspectCmd = &cobra.Command{
	Use:   "inspect [options] NETWORK [NETWORK...]",
	Short: "Inspect network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_inspectCmd).Standalone()

	network_inspectCmd.Flags().StringP("format", "f", "", "Pretty-print network to JSON or using a Go template")
	networkCmd.AddCommand(network_inspectCmd)
}
