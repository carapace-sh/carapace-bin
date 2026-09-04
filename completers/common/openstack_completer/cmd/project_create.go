package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_createCmd).Standalone()

	project_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	project_createCmd.Flags().String("description", "", "Project description")
	project_createCmd.Flags().Bool("disable", false, "Disable project")
	project_createCmd.Flags().String("domain", "", "Domain owning the project (name or ID)")
	project_createCmd.Flags().Bool("enable", false, "Enable project")
	project_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	project_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	project_createCmd.Flags().Bool("immutable", false, "Make resource immutable.")
	project_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	project_createCmd.Flags().Bool("no-immutable", false, "Make resource mutable (default)")
	project_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	project_createCmd.Flags().Bool("or-show", false, "Return existing project")
	project_createCmd.Flags().String("parent", "", "Parent of the project (name or ID)")
	project_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	project_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	project_createCmd.Flags().String("property", "", "Add a property to <name> (repeat option to set multiple properties)")
	project_createCmd.Flags().String("tag", "", "Tag to be added to the project (repeat option to set multiple tags)")
	project_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	projectCmd.AddCommand(project_createCmd)
}
