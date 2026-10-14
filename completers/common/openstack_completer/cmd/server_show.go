package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show server details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_showCmd).Standalone()

	server_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_showCmd.Flags().Bool("diagnostics", false, "Display server diagnostics information")
	server_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_showCmd.Flags().Bool("topology", false, "Include topology information in the output (supported by --os-compute-api-version 2.78 or above)")
	server_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	serverCmd.AddCommand(server_showCmd)
}
