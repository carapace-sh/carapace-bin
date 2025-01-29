package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats [options] [CONTAINER...]",
	Short: "Display a live stream of container resource usage statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	statsCmd.Flags().BoolP("all", "a", false, "Show all containers. Only running containers are shown by default. The default is false")
	statsCmd.Flags().String("format", "", "Pretty-print container statistics to JSON or using a Go template")
	statsCmd.Flags().StringP("interval", "i", "", "Time in seconds between stats reports")
	statsCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	statsCmd.Flags().Bool("no-reset", false, "Disable resetting the screen between intervals")
	statsCmd.Flags().Bool("no-stream", false, "Disable streaming stats and only pull the first result, default setting is false")
	statsCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	rootCmd.AddCommand(statsCmd)
}
