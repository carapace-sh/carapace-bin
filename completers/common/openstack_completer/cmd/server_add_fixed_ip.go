package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_add_fixed_ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Add fixed IP address to server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_add_fixed_ipCmd).Standalone()

	server_add_fixed_ipCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_add_fixed_ipCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_add_fixed_ipCmd.Flags().String("fixed-ip-address", "", "Requested fixed IP address")
	server_add_fixed_ipCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_add_fixed_ipCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_add_fixed_ipCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_add_fixed_ipCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	server_add_fixed_ipCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_add_fixed_ipCmd.Flags().String("tag", "", "Tag for the attached interface.")
	server_add_fixed_ipCmd.Flags().String("variable", "", "==SUPPRESS==")
	server_add_fixedCmd.AddCommand(server_add_fixed_ipCmd)
}
