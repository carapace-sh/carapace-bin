package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var node_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] self|NODE [NODE...]",
	Short: "Display detailed information on one or more nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_inspectCmd).Standalone()

	node_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	node_inspectCmd.Flags().Bool("pretty", false, "Print the information in a human friendly format")
	nodeCmd.AddCommand(node_inspectCmd)

	carapace.Gen(node_inspectCmd).PositionalAnyCompletion(docker.ActionNodes())
}
