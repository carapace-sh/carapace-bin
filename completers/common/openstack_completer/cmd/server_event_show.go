package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_event_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show server event details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_event_showCmd).Standalone()

	server_event_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_event_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_event_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_event_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_event_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_event_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_event_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_event_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_eventCmd.AddCommand(server_event_showCmd)
}
