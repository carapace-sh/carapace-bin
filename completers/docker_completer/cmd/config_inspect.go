package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var config_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] CONFIG [CONFIG...]",
	Short: "Display detailed information on one or more configs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_inspectCmd).Standalone()

	config_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	config_inspectCmd.Flags().Bool("pretty", false, "Print the information in a human friendly format")
	configCmd.AddCommand(config_inspectCmd)

	carapace.Gen(config_inspectCmd).PositionalAnyCompletion(docker.ActionConfigs())
}
