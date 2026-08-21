package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var role_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(role_createCmd).Standalone()

	role_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	role_createCmd.Flags().String("description", "", "Add description about the role")
	role_createCmd.Flags().String("domain", "", "Domain the role belongs to (name or ID)")
	role_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	role_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	role_createCmd.Flags().Bool("immutable", false, "Make resource immutable.")
	role_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	role_createCmd.Flags().Bool("no-immutable", false, "Make resource mutable (default)")
	role_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	role_createCmd.Flags().Bool("or-show", false, "Return existing role")
	role_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	role_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	role_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	roleCmd.AddCommand(role_createCmd)
}
