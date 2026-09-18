package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_scope_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Address Scope",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_scope_createCmd).Standalone()

	address_scope_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	address_scope_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	address_scope_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	address_scope_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	address_scope_createCmd.Flags().String("ip-version", "", "IP version (default is 4)")
	address_scope_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	address_scope_createCmd.Flags().Bool("no-share", false, "Do not share the address scope between projects (default)")
	address_scope_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	address_scope_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	address_scope_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	address_scope_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	address_scope_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	address_scope_createCmd.Flags().Bool("share", false, "Share the address scope between projects")
	address_scope_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	address_scopeCmd.AddCommand(address_scope_createCmd)
}
