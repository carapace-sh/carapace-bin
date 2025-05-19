package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var node_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List nodes in the swarm",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_lsCmd).Standalone()

	node_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	node_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	node_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	nodeCmd.AddCommand(node_lsCmd)
}
