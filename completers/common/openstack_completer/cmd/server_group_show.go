package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_group_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display server group details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_group_showCmd).Standalone()

	server_group_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_group_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_group_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_group_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_group_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_group_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_group_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_group_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_groupCmd.AddCommand(server_group_showCmd)
}
