package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var volume_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] VOLUME [VOLUME...]",
	Short: "Display detailed information on one or more volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_inspectCmd).Standalone()

	volume_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	volumeCmd.AddCommand(volume_inspectCmd)

	carapace.Gen(volume_inspectCmd).PositionalAnyCompletion(docker.ActionVolumes())
}
