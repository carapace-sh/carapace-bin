package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_type_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display share type details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_type_showCmd).Standalone()

	share_type_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_type_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_type_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_type_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_type_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_type_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_type_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_type_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_typeCmd.AddCommand(share_type_showCmd)
}
