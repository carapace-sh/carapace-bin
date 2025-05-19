package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List networks",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_lsCmd).Standalone()

	network_lsCmd.Flags().StringP("filter", "f", "", "Provide filter values (e.g. \"driver=bridge\")")
	network_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	network_lsCmd.Flags().Bool("no-trunc", false, "Do not truncate the output")
	network_lsCmd.Flags().BoolP("quiet", "q", false, "Only display network IDs")
	networkCmd.AddCommand(network_lsCmd)
}
