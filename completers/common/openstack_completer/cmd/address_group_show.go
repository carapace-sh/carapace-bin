package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_group_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display address group details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_group_showCmd).Standalone()

	address_group_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	address_group_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	address_group_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	address_group_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	address_group_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	address_group_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	address_group_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	address_group_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	address_groupCmd.AddCommand(address_group_showCmd)
}
