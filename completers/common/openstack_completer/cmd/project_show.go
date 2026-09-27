package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display project details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_showCmd).Standalone()

	project_showCmd.Flags().Bool("children", false, "Show project's subtree (children) as a list")
	project_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	project_showCmd.Flags().String("domain", "", "Domain owning <project> (name or ID)")
	project_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	project_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	project_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	project_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	project_showCmd.Flags().Bool("parents", false, "Show the project's parents as a list")
	project_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	project_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	project_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	projectCmd.AddCommand(project_showCmd)
}
