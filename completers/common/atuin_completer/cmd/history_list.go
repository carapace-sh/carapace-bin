package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items in history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_listCmd).Standalone()

	history_listCmd.Flags().Bool("cmd-only", false, "Show only the text of the command")
	history_listCmd.Flags().BoolP("cwd", "c", false, "")
	history_listCmd.Flags().StringP("format", "f", "", "Available variables: {command}, {directory}, {duration}, {user}, {host}, {exit}, {time}, {session}, and {uuid} Example: --format \"{time} - [{duration}] - {directory}$\\t{command}\"")
	history_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	history_listCmd.Flags().Bool("human", false, "")
	history_listCmd.Flags().Bool("print0", false, "Terminate the output with a null, for better multiline support")
	history_listCmd.Flags().StringP("reverse", "r", "", "")
	history_listCmd.Flags().BoolP("session", "s", false, "")
	history_listCmd.Flags().String("timezone", "", "Display the command time in another timezone other than the configured default")
	history_listCmd.Flags().String("tz", "", "Display the command time in another timezone other than the configured default")
	historyCmd.AddCommand(history_listCmd)
}
