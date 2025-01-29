package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_psCmd = &cobra.Command{
	Use:   "ps [options]",
	Short: "List containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_psCmd).Standalone()

	container_psCmd.Flags().BoolP("all", "a", false, "Show all the containers, default is only running containers")
	container_psCmd.Flags().Bool("external", false, "Show containers in storage not controlled by Podman")
	container_psCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	container_psCmd.Flags().String("format", "", "Pretty-print containers to JSON or using a Go template")
	container_psCmd.Flags().StringP("last", "n", "", "Print the n last created containers (all states)")
	container_psCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_psCmd.Flags().Bool("no-trunc", false, "Display the extended information")
	container_psCmd.Flags().Bool("noheading", false, "Do not print headers")
	container_psCmd.Flags().Bool("ns", false, "Display namespace information")
	container_psCmd.Flags().BoolP("pod", "p", false, "Print the ID and name of the pod the containers are associated with")
	container_psCmd.Flags().BoolP("quiet", "q", false, "Print the numeric IDs of the containers only")
	container_psCmd.Flags().BoolP("size", "s", false, "Display the total file sizes")
	container_psCmd.Flags().String("sort", "", "Sort output by: command, created, id, image, names, runningfor, size, status")
	container_psCmd.Flags().Bool("sync", false, "Sync container state with OCI runtime")
	container_psCmd.Flags().StringP("watch", "w", "", "Watch the ps output on an interval in seconds")
	containerCmd.AddCommand(container_psCmd)
}
