package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Get the last command ran",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_lastCmd).Standalone()

	history_lastCmd.Flags().Bool("cmd-only", false, "Show only the text of the command")
	history_lastCmd.Flags().StringP("format", "f", "", "Available variables: {command}, {directory}, {duration}, {user}, {host}, {time}, {session}, {uuid} and {relativetime}. Example: --format \"{time} - [{duration}] - {directory}$\\t{command}\"")
	history_lastCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_lastCmd.Flags().Bool("human", false, "")
	history_lastCmd.Flags().String("timezone", "", "Display the command time in another timezone other than the configured default")
	history_lastCmd.Flags().String("tz", "", "Display the command time in another timezone other than the configured default")
	historyCmd.AddCommand(history_lastCmd)
}
