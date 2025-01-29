package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_listCmd = &cobra.Command{
	Use:     "list [options]",
	Short:   "List containers",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_listCmd).Standalone()

	container_listCmd.Flags().BoolP("all", "a", false, "Show all the containers, default is only running containers")
	container_listCmd.Flags().Bool("external", false, "Show containers in storage not controlled by Podman")
	container_listCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions given")
	container_listCmd.Flags().String("format", "", "Pretty-print containers to JSON or using a Go template")
	container_listCmd.Flags().StringP("last", "n", "", "Print the n last created containers (all states)")
	container_listCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_listCmd.Flags().Bool("no-trunc", false, "Display the extended information")
	container_listCmd.Flags().Bool("noheading", false, "Do not print headers")
	container_listCmd.Flags().Bool("ns", false, "Display namespace information")
	container_listCmd.Flags().BoolP("pod", "p", false, "Print the ID and name of the pod the containers are associated with")
	container_listCmd.Flags().BoolP("quiet", "q", false, "Print the numeric IDs of the containers only")
	container_listCmd.Flags().BoolP("size", "s", false, "Display the total file sizes")
	container_listCmd.Flags().String("sort", "", "Sort output by: command, created, id, image, names, runningfor, size, status")
	container_listCmd.Flags().Bool("sync", false, "Sync container state with OCI runtime")
	container_listCmd.Flags().StringP("watch", "w", "", "Watch the ps output on an interval in seconds")
	containerCmd.AddCommand(container_listCmd)
}
