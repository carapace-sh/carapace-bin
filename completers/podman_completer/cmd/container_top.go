package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_topCmd = &cobra.Command{
	Use:   "top [options] CONTAINER [FORMAT-DESCRIPTORS|ARGS...]",
	Short: "Display the running processes of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_topCmd).Standalone()

	container_topCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_topCmd.Flags().Bool("list-descriptors", false, "")
	container_topCmd.Flag("list-descriptors").Hidden = true
	containerCmd.AddCommand(container_topCmd)
}
