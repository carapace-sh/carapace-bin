package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Add a share to a server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_shareCmd).Standalone()

	server_add_shareCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_add_shareCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_add_shareCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_add_shareCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_add_shareCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_add_shareCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_add_shareCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_add_shareCmd.Flags().String("tag", "", "Optional tag used to mount the share, if not provided the share uuid is used as tag by default")
	server_add_shareCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_addCmd.AddCommand(server_add_shareCmd)
}
