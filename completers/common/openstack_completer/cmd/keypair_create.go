package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keypair_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new public or private key for server ssh access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keypair_createCmd).Standalone()

	keypair_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	keypair_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	keypair_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	keypair_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	keypair_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	keypair_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	keypair_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	keypair_createCmd.Flags().String("private-key", "", "Filename for private key to save.")
	keypair_createCmd.Flags().String("public-key", "", "Filename for public key to add.")
	keypair_createCmd.Flags().String("type", "", "Keypair type (supported by --os-compute-api-version 2.2 or above)")
	keypair_createCmd.Flags().String("user", "", "The owner of the keypair (admin only) (name or ID) (supported by --os-compute-api-version 2.10 or above)")
	keypair_createCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	keypair_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	keypairCmd.AddCommand(keypair_createCmd)
}
