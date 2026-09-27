package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var credential_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new credential",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credential_createCmd).Standalone()

	credential_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	credential_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	credential_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	credential_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	credential_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	credential_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	credential_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	credential_createCmd.Flags().String("project", "", "Project which limits the scope of the credential (name or ID)")
	credential_createCmd.Flags().String("type", "", "New credential type: cert, ec2, totp and so on")
	credential_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	credentialCmd.AddCommand(credential_createCmd)
}
