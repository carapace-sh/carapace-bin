package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs [options] CONTAINER [CONTAINER...]",
	Short: "Fetch the logs of one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().Bool("color", false, "Output the containers with different colors in the log.")
	logsCmd.Flags().Bool("details", false, "Show extra details provided to the logs")
	logsCmd.Flags().BoolP("follow", "f", false, "Follow log output.  The default is false")
	logsCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	logsCmd.Flags().BoolP("names", "n", false, "Output the container name in the log")
	logsCmd.Flags().String("since", "", "Show logs since TIMESTAMP")
	logsCmd.Flags().String("tail", "", "Output the specified number of LINES at the end of the logs.  Defaults to -1, which prints all lines")
	logsCmd.Flags().BoolP("timestamps", "t", false, "Output the timestamps in the log")
	logsCmd.Flags().String("until", "", "Show logs until TIMESTAMP")
	logsCmd.Flag("details").Hidden = true
	rootCmd.AddCommand(logsCmd)
}
