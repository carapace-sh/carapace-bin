package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "drift",
	Short: "A standalone git diff pager for the terminal",
	Long:  "https://github.com/aymanbagabas/uncurses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "A", false, "Also show untracked (non-ignored) files as new-file diffs. Only affects the working-tree view")
	rootCmd.Flags().Bool("cached", false, "Show staged changes (index vs HEAD)")
	rootCmd.Flags().StringP("config", "c", "", "Path to a config file (toml/yaml/json). Overrides auto-discovery")
	rootCmd.Flags().StringP("context", "U", "", "Lines of context around each change (git `-U`)")
	rootCmd.Flags().String("diff-algorithm", "", "Diff algorithm: myers, minimal, patience, or histogram")
	rootCmd.Flags().StringP("directory", "C", "", "Run as if drift was started in this directory (like `git -C`). Works with worktrees and bare repositories")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().Bool("ignore-whitespace", false, "Ignore whitespace-only changes (git `-w`)")
	rootCmd.Flags().String("interval", "300", "How often (ms) watch mode polls for unstaged working-tree edits, which the git index/refs watcher can't see")
	rootCmd.Flags().Bool("no-syntax", false, "Disable syntax highlighting for this run")
	rootCmd.Flags().Bool("staged", false, "Show staged changes (index vs HEAD)")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().BoolP("watch", "w", false, "Watch the git index and refs and refresh the diff on change")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":         carapace.ActionFiles(),
		"diff-algorithm": carapace.ActionValues("myers", "minimal", "patience", "histogram"),
		"directory":      carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, _ *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(rootCmd.Flag("directory")))
	})
}
