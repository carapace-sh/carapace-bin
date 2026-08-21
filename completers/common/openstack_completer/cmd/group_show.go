package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display group details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_showCmd).Standalone()

	group_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	group_showCmd.Flags().String("domain", "", "Domain containing <group> (name or ID)")
	group_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	group_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	group_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	group_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	group_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	group_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	group_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	groupCmd.AddCommand(group_showCmd)
}
