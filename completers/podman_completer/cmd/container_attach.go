package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_attachCmd = &cobra.Command{
	Use:   "attach [options] CONTAINER",
	Short: "Attach to a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_attachCmd).Standalone()

	container_attachCmd.Flags().String("detach-keys", "", "Select the key sequence for detaching a container")
	container_attachCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_attachCmd.Flags().Bool("no-stdin", false, "Do not attach STDIN. The default is false")
	container_attachCmd.Flags().Bool("sig-proxy", false, "Proxy received signals to the process")
	containerCmd.AddCommand(container_attachCmd)
}
