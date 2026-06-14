package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var workspace_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_listCmd).Standalone()

	workspace_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspace_listCmd.Flags().StringP("template", "T", "", "Render each workspace using the given template")
	workspaceCmd.AddCommand(workspace_listCmd)

	carapace.Gen(workspace_listCmd).FlagCompletion(carapace.ActionMap{
		"template": jj.ActionTemplates(),
	})
}
