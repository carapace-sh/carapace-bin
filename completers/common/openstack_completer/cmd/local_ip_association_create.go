package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var local_ip_association_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Local IP Association",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(local_ip_association_createCmd).Standalone()

	local_ip_association_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	local_ip_association_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	local_ip_association_createCmd.Flags().String("fixed-ip", "", "Fixed IP for Local IP Association")
	local_ip_association_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	local_ip_association_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	local_ip_association_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	local_ip_association_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	local_ip_association_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	local_ip_association_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	local_ip_association_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	local_ip_associationCmd.AddCommand(local_ip_association_createCmd)
}
