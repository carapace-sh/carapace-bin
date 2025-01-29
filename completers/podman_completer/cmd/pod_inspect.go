package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_inspectCmd = &cobra.Command{
	Use:   "inspect [options] POD [POD...]",
	Short: "Display a pod configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_inspectCmd).Standalone()

	pod_inspectCmd.Flags().StringP("format", "f", "", "Format the output to a Go template or json")
	pod_inspectCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	podCmd.AddCommand(pod_inspectCmd)
}
