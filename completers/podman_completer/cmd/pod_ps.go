package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_psCmd = &cobra.Command{
	Use:     "ps [options]",
	Short:   "List pods",
	Aliases: []string{"ls", "list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_psCmd).Standalone()

	pod_psCmd.Flags().Bool("ctr-ids", false, "Display the container UUIDs. If no-trunc is not set they will be truncated")
	pod_psCmd.Flags().Bool("ctr-names", false, "Display the container names")
	pod_psCmd.Flags().Bool("ctr-status", false, "Display the container status")
	pod_psCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	pod_psCmd.Flags().String("format", "", "Pretty-print pods to JSON or using a Go template")
	pod_psCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	pod_psCmd.Flags().Bool("no-trunc", false, "Do not truncate pod and container IDs")
	pod_psCmd.Flags().BoolP("noheading", "n", false, "Do not print headers")
	pod_psCmd.Flags().Bool("ns", false, "Display namespace information of the pod")
	pod_psCmd.Flags().BoolP("quiet", "q", false, "Print the numeric IDs of the pods only")
	pod_psCmd.Flags().String("sort", "", "Sort output by created, id, name, or number")
	podCmd.AddCommand(pod_psCmd)
}
