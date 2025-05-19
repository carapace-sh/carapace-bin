package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergetoolCmd = &cobra.Command{
	Use:     "mergetool",
	Short:   "Run merge conflict resolution tools to resolve merge conflicts",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_manipulator].ID,
}

func init() {
	carapace.Gen(mergetoolCmd).Standalone()

	mergetoolCmd.Flags().BoolP("gui", "g", false, "Read from the configured merge.guitool variable instead of merge.tool.")
	mergetoolCmd.Flags().Bool("no-gui", false, "This overrides a previous -g or --gui setting")
	mergetoolCmd.Flags().BoolP("no-prompt", "y", false, "Don’t prompt before each invocation of the merge resolution program")
	mergetoolCmd.Flags().Bool("prompt", false, "Prompt before each invocation of the merge resolution program to give the user a chance to skip the path")
	mergetoolCmd.Flags().String("tool", "", "Use the merge resolution program specified by <tool>")
	mergetoolCmd.Flags().Bool("tool-help", false, "Print a list of merge tools that may be used with --tool")
	rootCmd.AddCommand(mergetoolCmd)

	mergetoolCmd.Flag("tool").NoOptDefVal = " "

	carapace.Gen(mergetoolCmd).FlagCompletion(carapace.ActionMap{
		"tool": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(mergetoolCmd).PositionalAnyCompletion(
		git.ActionUnmergedFiles(),
	)
}
