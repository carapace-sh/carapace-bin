package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domain_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domain_createCmd).Standalone()

	domain_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	domain_createCmd.Flags().String("description", "", "New domain description")
	domain_createCmd.Flags().Bool("disable", false, "Disable domain")
	domain_createCmd.Flags().Bool("enable", false, "Enable domain (default)")
	domain_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	domain_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	domain_createCmd.Flags().Bool("immutable", false, "Make resource immutable.")
	domain_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	domain_createCmd.Flags().Bool("no-immutable", false, "Make resource mutable (default)")
	domain_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	domain_createCmd.Flags().Bool("or-show", false, "Return existing domain")
	domain_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	domain_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	domain_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	domainCmd.AddCommand(domain_createCmd)
}
