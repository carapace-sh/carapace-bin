package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_lsCmd).Standalone()

	container_lsCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	container_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	container_lsCmd.Flags().String("format", "", "Pretty-print containers using a Go template")
	container_lsCmd.Flags().StringP("last", "n", "", "Show n last created containers (includes all states) (default -1)")
	container_lsCmd.Flags().BoolP("latest", "l", false, "Show the latest created container (includes all states)")
	container_lsCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	container_lsCmd.Flags().BoolP("quiet", "q", false, "Only display numeric IDs")
	container_lsCmd.Flags().BoolP("size", "s", false, "Display total file sizes")
	containerCmd.AddCommand(container_lsCmd)

	carapace.Gen(container_lsCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionValues("ancestor=", "exited=", "health=", "label=", "network=", "since=", "volume=", "before=", "expose=", "id=", "name=", "publish=", "status="),
	})
}
