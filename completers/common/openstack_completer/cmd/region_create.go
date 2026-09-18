package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var region_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new region",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(region_createCmd).Standalone()

	region_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	region_createCmd.Flags().String("description", "", "New region description")
	region_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	region_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	region_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	region_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	region_createCmd.Flags().String("parent-region", "", "Parent region ID")
	region_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	region_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	region_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	regionCmd.AddCommand(region_createCmd)
}
