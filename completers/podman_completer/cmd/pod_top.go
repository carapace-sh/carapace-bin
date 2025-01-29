package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_topCmd = &cobra.Command{
	Use:   "top [options] POD [FORMAT-DESCRIPTORS|ARGS...]",
	Short: "Display the running processes of containers in a pod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_topCmd).Standalone()

	pod_topCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	pod_topCmd.Flags().Bool("list-descriptors", false, "")
	pod_topCmd.Flag("list-descriptors").Hidden = true
	podCmd.AddCommand(pod_topCmd)
}
