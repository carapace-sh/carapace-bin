package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var role_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display role details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(role_showCmd).Standalone()

	role_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	role_showCmd.Flags().String("domain", "", "Domain the role belongs to (name or ID)")
	role_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	role_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	role_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	role_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	role_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	role_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	role_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	roleCmd.AddCommand(role_showCmd)
}
