package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_diffCmd = &cobra.Command{
	Use:   "diff [options] CONTAINER [CONTAINER]",
	Short: "Inspect changes to the container's file systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_diffCmd).Standalone()

	container_diffCmd.Flags().String("format", "", "Change the output format (json)")
	container_diffCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	containerCmd.AddCommand(container_diffCmd)
}
