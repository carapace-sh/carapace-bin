package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats [OPTIONS] [SERVICE]",
	Short: "Display a live stream of container(s) resource usage statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	statsCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	statsCmd.Flags().String("format", "", "Format output using a custom template:")
	statsCmd.Flags().Bool("no-stream", false, "Disable streaming stats and only pull the first result")
	statsCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	rootCmd.AddCommand(statsCmd)

	carapace.Gen(statsCmd).PositionalCompletion(
		action.ActionServices(statsCmd),
	)
}
