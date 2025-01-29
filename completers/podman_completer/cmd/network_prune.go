package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_pruneCmd = &cobra.Command{
	Use:   "prune [options]",
	Short: "Prune unused networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_pruneCmd).Standalone()

	network_pruneCmd.Flags().StringSlice("filter", []string{}, "Provide filter values (e.g. 'label=<key>=<value>')")
	network_pruneCmd.Flags().BoolP("force", "f", false, "do not prompt for confirmation")
	networkCmd.AddCommand(network_pruneCmd)
}
