package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_share_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show detail of a share attachment to a server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_share_showCmd).Standalone()

	server_share_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_share_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_share_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_share_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_share_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_share_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_share_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_share_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_shareCmd.AddCommand(server_share_showCmd)
}
