package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_logsCmd = &cobra.Command{
	Use:   "logs [OPTIONS] CONTAINER",
	Short: "Fetch the logs of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_logsCmd).Standalone()

	container_logsCmd.Flags().Bool("details", false, "Show extra details provided to logs")
	container_logsCmd.Flags().BoolP("follow", "f", false, "Follow log output")
	container_logsCmd.Flags().String("since", "", "Show logs since timestamp (e.g. \"2013-01-02T13:23:37Z\") or relative (e.g. \"42m\" for 42 minutes)")
	container_logsCmd.Flags().StringP("tail", "n", "all", "Number of lines to show from the end of the logs")
	container_logsCmd.Flags().BoolP("timestamps", "t", false, "Show timestamps")
	container_logsCmd.Flags().String("until", "", "Show logs before a timestamp (e.g. \"2013-01-02T13:23:37Z\") or relative (e.g. \"42m\" for 42 minutes)")
	containerCmd.AddCommand(container_logsCmd)

	carapace.Gen(container_logsCmd).PositionalCompletion(
		docker.ActionContainers(),
	)
}
