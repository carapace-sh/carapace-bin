package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] SERVICE [SERVICE...]",
	Short: "Display detailed information on one or more services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_inspectCmd).Standalone()

	service_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	service_inspectCmd.Flags().Bool("pretty", false, "Print the information in a human friendly format")
	serviceCmd.AddCommand(service_inspectCmd)

	carapace.Gen(service_inspectCmd).PositionalAnyCompletion(docker.ActionServices())
}
