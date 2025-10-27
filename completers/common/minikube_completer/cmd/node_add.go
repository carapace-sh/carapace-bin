package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var node_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a node to the given cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_addCmd).Standalone()

	node_addCmd.Flags().Bool("control-plane", false, "If set, added node will become a control-plane. Defaults to false. Currently only supported for existing HA (multi-control plane) clusters.")
	node_addCmd.Flags().Bool("delete-on-failure", false, "If set, delete the current cluster if start fails and try again. Defaults to false.")
	node_addCmd.Flags().Bool("worker", false, "If set, added node will be available as worker. Defaults to true.")
	nodeCmd.AddCommand(node_addCmd)
}
