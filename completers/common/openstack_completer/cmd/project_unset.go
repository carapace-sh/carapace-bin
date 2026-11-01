package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var project_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset project properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_unsetCmd).Standalone()

	project_unsetCmd.Flags().String("property", "", "Unset a project property (repeat option to unset multiple properties)")
	projectCmd.AddCommand(project_unsetCmd)

	carapace.Gen(project_unsetCmd).PositionalCompletion(action.ActionProjects(project_unsetCmd))
}
