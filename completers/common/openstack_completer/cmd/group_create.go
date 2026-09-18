package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_createCmd).Standalone()

	group_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	group_createCmd.Flags().String("description", "", "New group description")
	group_createCmd.Flags().String("domain", "", "Domain to contain new group (name or ID)")
	group_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	group_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	group_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	group_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	group_createCmd.Flags().Bool("or-show", false, "Return existing group")
	group_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	group_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	group_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	groupCmd.AddCommand(group_createCmd)
}
