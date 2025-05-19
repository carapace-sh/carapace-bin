package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/spf13/cobra"
)

var difftoolCmd = &cobra.Command{
	Use:     "difftool",
	Short:   "Show changes using common diff tools",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(difftoolCmd).Standalone()

	difftoolCmd.Flags().BoolP("dir-diff", "d", false, "perform a full-directory diff")
	difftoolCmd.Flags().StringP("extcmd", "x", "", "specify a custom command for viewing diffs")
	difftoolCmd.Flags().BoolP("gui", "g", false, "use `diff.guitool` instead of `diff.tool`")
	difftoolCmd.Flags().BoolP("no-prompt", "y", false, "do not prompt before launching a diff tool")
	difftoolCmd.Flags().Bool("symlinks", false, "use symlinks in dir-diff mode")
	difftoolCmd.Flags().StringP("tool", "t", "", "use the specified diff tool")
	difftoolCmd.Flags().Bool("tool-help", false, "print a list of diff tools that may be used with `--tool`")
	difftoolCmd.Flags().Bool("trust-exit-code", false, "exit when an invoked diff tool returns a non - zero exit code")

	common.AddDiffFlags(difftoolCmd)
	rootCmd.AddCommand(difftoolCmd)

	carapace.Gen(difftoolCmd).FlagCompletion(carapace.ActionMap{
		"tool": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(difftoolCmd).PositionalAnyCompletion(
		actionDiffArgs(difftoolCmd),
	)

	carapace.Gen(difftoolCmd).DashAnyCompletion(
		carapace.ActionPositional(difftoolCmd),
	)
}
