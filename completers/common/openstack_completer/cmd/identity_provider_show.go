package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identity_provider_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display identity provider details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identity_provider_showCmd).Standalone()

	identity_provider_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	identity_provider_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	identity_provider_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	identity_provider_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	identity_provider_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	identity_provider_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	identity_provider_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	identity_provider_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	identity_providerCmd.AddCommand(identity_provider_showCmd)
}
