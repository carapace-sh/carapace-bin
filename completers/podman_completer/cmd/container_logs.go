package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_logsCmd = &cobra.Command{
	Use:   "logs [options] CONTAINER [CONTAINER...]",
	Short: "Fetch the logs of one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_logsCmd).Standalone()

	container_logsCmd.Flags().Bool("color", false, "Output the containers with different colors in the log.")
	container_logsCmd.Flags().Bool("details", false, "Show extra details provided to the logs")
	container_logsCmd.Flags().BoolP("follow", "f", false, "Follow log output.  The default is false")
	container_logsCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_logsCmd.Flags().BoolP("names", "n", false, "Output the container name in the log")
	container_logsCmd.Flags().String("since", "", "Show logs since TIMESTAMP")
	container_logsCmd.Flags().String("tail", "", "Output the specified number of LINES at the end of the logs.  Defaults to -1, which prints all lines")
	container_logsCmd.Flags().BoolP("timestamps", "t", false, "Output the timestamps in the log")
	container_logsCmd.Flags().String("until", "", "Show logs until TIMESTAMP")
	container_logsCmd.Flag("details").Hidden = true
	containerCmd.AddCommand(container_logsCmd)
}
