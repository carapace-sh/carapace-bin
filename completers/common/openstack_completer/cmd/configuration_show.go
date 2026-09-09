package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configuration_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display configuration details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configuration_showCmd).Standalone()

	configuration_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	configuration_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	configuration_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	configuration_showCmd.Flags().Bool("mask", false, "Attempt to mask passwords (default)")
	configuration_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	configuration_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	configuration_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	configuration_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	configuration_showCmd.Flags().Bool("unmask", false, "Show password in clear text")
	configuration_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	configurationCmd.AddCommand(configuration_showCmd)
}
