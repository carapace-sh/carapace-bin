package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var workspaces_foreachCmd = &cobra.Command{
	Use:   "foreach",
	Short: "Run a command on all workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_foreachCmd).Standalone()

	workspaces_foreachCmd.Flags().BoolP("all", "A", false, "Run the command on all workspaces of a project")
	workspaces_foreachCmd.Flags().String("exclude", "", "An array of glob pattern idents; matching workspaces won't be traversed")
	workspaces_foreachCmd.Flags().String("from", "", "An array of glob pattern idents from which to base any recursion")
	workspaces_foreachCmd.Flags().String("include", "", "An array of glob pattern idents; only matching workspaces will be traversed")
	workspaces_foreachCmd.Flags().BoolP("interlaced", "i", false, "Print the output of commands in real-time instead of buffering it")
	workspaces_foreachCmd.Flags().StringP("jobs", "j", "", "The maximum number of parallel tasks that the execution will be limited to; or unlimited")
	workspaces_foreachCmd.Flags().Bool("no-private", false, "Avoid running the command on private workspaces")
	workspaces_foreachCmd.Flags().BoolP("parallel", "p", false, "Run the commands in parallel")
	workspaces_foreachCmd.Flags().BoolP("recursive", "R", false, "Find packages via dependencies/devDependencies instead of using the workspaces field")
	workspaces_foreachCmd.Flags().Bool("since", false, "Only include workspaces that have been changed since the specified ref.")
	workspaces_foreachCmd.Flags().BoolP("topological", "t", false, "Run the command after all workspaces it depends on (regular) have finished")
	workspaces_foreachCmd.Flags().Bool("topological-dev", false, "Run the command after all workspaces it depends on (regular + dev) have finished")
	workspaces_foreachCmd.Flags().BoolP("verbose", "v", false, "Prefix each output line with the name of the originating workspace")
	workspacesCmd.AddCommand(workspaces_foreachCmd)

	carapace.Gen(workspaces_foreachCmd).FlagCompletion(carapace.ActionMap{
		"exclude": yarn.ActionWorkspaces().UniqueList(","),
		"include": yarn.ActionWorkspaces().UniqueList(","),
	})

	// TODO positional completion
}
