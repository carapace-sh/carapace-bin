package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var user_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display user details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_showCmd).Standalone()

	user_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	user_showCmd.Flags().String("domain", "", "Domain owning <user> (name or ID)")
	user_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	user_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	user_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	user_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	user_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	user_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	user_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	userCmd.AddCommand(user_showCmd)
}
