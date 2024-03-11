package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_listCmd).Standalone()

	workspaces_listCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	workspaces_listCmd.Flags().Bool("no-private", false, "Exclude workspaces that have the private field set to true")
	workspaces_listCmd.Flags().BoolP("recursive", "R", false, "Find packages via dependencies/devDependencies instead of using the workspaces field")
	workspaces_listCmd.Flags().Bool("since", false, "Only include workspaces that have been changed since the specified ref.")
	workspaces_listCmd.Flags().BoolP("verbose", "v", false, "Also return the cross-dependencies between workspaces")
	workspacesCmd.AddCommand(workspaces_listCmd)
}
