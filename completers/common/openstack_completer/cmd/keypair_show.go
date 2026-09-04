package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keypair_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display key details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keypair_showCmd).Standalone()

	keypair_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	keypair_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	keypair_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	keypair_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	keypair_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	keypair_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	keypair_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	keypair_showCmd.Flags().Bool("public-key", false, "Show only bare public key paired with the generated key")
	keypair_showCmd.Flags().String("user", "", "The owner of the keypair.")
	keypair_showCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	keypair_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	keypairCmd.AddCommand(keypair_showCmd)
}
