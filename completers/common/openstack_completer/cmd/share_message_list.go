package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_message_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all messages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_message_listCmd).Standalone()

	share_message_listCmd.Flags().String("action-id", "", "Filters results by action id.")
	share_message_listCmd.Flags().String("before", "", "Return only user messages created before given date.")
	share_message_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_message_listCmd.Flags().String("detail-id", "", "Filters results by detail id.")
	share_message_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_message_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_message_listCmd.Flags().String("limit", "", "Maximum number of messages to return.")
	share_message_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_message_listCmd.Flags().String("message-level", "", "Filters results by the message level.")
	share_message_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_message_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_message_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_message_listCmd.Flags().String("request-id", "", "Filters results by request id.")
	share_message_listCmd.Flags().String("resource-id", "", "Filters results by a resource uuid.")
	share_message_listCmd.Flags().String("resource-type", "", "Filters results by a resource type.")
	share_message_listCmd.Flags().String("since", "", "Return only user messages created since given date.")
	share_message_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_message_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_message_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_messageCmd.AddCommand(share_message_listCmd)
}
