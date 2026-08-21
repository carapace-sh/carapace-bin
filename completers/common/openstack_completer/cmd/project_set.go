package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set project properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_setCmd).Standalone()

	project_setCmd.Flags().Bool("clear-tags", false, "Clear tags associated with the project.")
	project_setCmd.Flags().String("description", "", "Set project description")
	project_setCmd.Flags().Bool("disable", false, "Disable project")
	project_setCmd.Flags().String("domain", "", "Domain owning <project> (name or ID)")
	project_setCmd.Flags().Bool("enable", false, "Enable project")
	project_setCmd.Flags().Bool("immutable", false, "Make resource immutable.")
	project_setCmd.Flags().String("name", "", "Set project name")
	project_setCmd.Flags().Bool("no-immutable", false, "Make resource mutable (default)")
	project_setCmd.Flags().String("property", "", "Set a property on <project> (repeat option to set multiple properties)")
	project_setCmd.Flags().String("remove-tag", "", "Tag to be deleted from the project (repeat option to delete multiple tags)")
	project_setCmd.Flags().String("tag", "", "Tag to be added to the project (repeat option to set multiple tags)")
	projectCmd.AddCommand(project_setCmd)
}
