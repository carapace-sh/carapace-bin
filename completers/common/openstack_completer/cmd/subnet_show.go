package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display subnet details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_showCmd).Standalone()

	subnet_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	subnet_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	subnet_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	subnet_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	subnet_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	subnet_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	subnet_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	subnet_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	subnetCmd.AddCommand(subnet_showCmd)
}
