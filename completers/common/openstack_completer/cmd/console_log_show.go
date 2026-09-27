package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var console_log_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show server's console output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(console_log_showCmd).Standalone()

	console_log_showCmd.Flags().String("lines", "", "Number of lines to display from the end of the log (default=all)")
	console_logCmd.AddCommand(console_log_showCmd)
}
