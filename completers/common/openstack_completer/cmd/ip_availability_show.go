package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ip_availability_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show network IP availability details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ip_availability_showCmd).Standalone()

	ip_availability_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	ip_availability_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	ip_availability_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	ip_availability_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	ip_availability_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	ip_availability_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	ip_availability_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	ip_availability_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	ip_availabilityCmd.AddCommand(ip_availability_showCmd)
}
