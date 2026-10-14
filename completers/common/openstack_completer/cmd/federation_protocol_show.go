package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_protocol_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display federation protocol details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_protocol_showCmd).Standalone()

	federation_protocol_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	federation_protocol_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	federation_protocol_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	federation_protocol_showCmd.Flags().String("identity-provider", "", "Identity provider that supports <federation-protocol> (name or ID) (required)")
	federation_protocol_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	federation_protocol_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	federation_protocol_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	federation_protocol_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	federation_protocol_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	federation_protocol_showCmd.MarkFlagRequired("identity-provider")
	federation_protocolCmd.AddCommand(federation_protocol_showCmd)
}
