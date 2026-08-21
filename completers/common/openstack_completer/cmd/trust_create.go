package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trust_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new trust",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_createCmd).Standalone()

	trust_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	trust_createCmd.Flags().String("expiration", "", "Sets an expiration date for the trust (format of YYYY-mm-ddTHH:MM:SS)")
	trust_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	trust_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	trust_createCmd.Flags().Bool("impersonate", false, "Tokens generated from the trust will represent <trustor> (defaults to False)")
	trust_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	trust_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	trust_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	trust_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	trust_createCmd.Flags().String("project", "", "Project being delegated (name or ID) (required)")
	trust_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	trust_createCmd.Flags().String("role", "", "Roles to authorize (name or ID) (repeat option to set multiple values, required)")
	trust_createCmd.Flags().String("trustee-domain", "", "Domain that contains <trustee> (name or ID)")
	trust_createCmd.Flags().String("trustor-domain", "", "Domain that contains <trustor> (name or ID)")
	trust_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	trust_createCmd.MarkFlagRequired("project")
	trust_createCmd.MarkFlagRequired("role")
	trustCmd.AddCommand(trust_createCmd)
}
