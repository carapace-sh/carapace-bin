package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var address_group_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Address Group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(address_group_createCmd).Standalone()

	address_group_createCmd.Flags().String("address", "", "IP address or CIDR (repeat option to set multiple addresses)")
	address_group_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	address_group_createCmd.Flags().String("description", "", "New address group description")
	address_group_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	address_group_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	address_group_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	address_group_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	address_group_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	address_group_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	address_group_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	address_group_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	address_group_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	address_group_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	address_groupCmd.AddCommand(address_group_createCmd)
}
