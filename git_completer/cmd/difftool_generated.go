package cmd

import (
	"github.com/spf13/cobra"
)

var difftoolCmd = &cobra.Command{
	Use: "difftool",
	Short: "Show changes using common diff tools",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
	difftoolCmd.Flags().BoolP("dir-diff", "d", false, "perform a full-directory diff")
	difftoolCmd.Flags().BoolP("gui", "g", false, "use `diff.guitool` instead of `diff.tool`")
	difftoolCmd.Flags().Bool("no-index", false, "passed to `diff`")
	difftoolCmd.Flags().Bool("symlinks", false, "use symlinks in dir-diff mode")
	difftoolCmd.Flags().Bool("tool-help", false, "print a list of diff tools that may be used with `--tool`")
	difftoolCmd.Flags().Bool("trust-exit-code", false, "make 'git-difftool' exit when an invoked diff tool returns a non - zero exit code")
	difftoolCmd.Flags().BoolP("tool", "t", false, "<tool>     use the specified diff tool")
	difftoolCmd.Flags().BoolP("extcmd", "x", false, "<command>    specify a custom command for viewing diffs")
	difftoolCmd.Flags().BoolP("no-prompt", "y", false, "do not prompt before launching a diff tool")
    rootCmd.AddCommand(difftoolCmd)
}
