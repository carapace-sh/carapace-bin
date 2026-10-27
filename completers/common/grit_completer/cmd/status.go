package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show what's changed and where you are (this is the default)",
	Aliases: []string{"st"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	statusCmd.Flags().Bool("ignored", false, "List ignored paths (traditional mode; same set as `git status --ignored`)")
	statusCmd.Flags().Bool("no-ahead-behind", false, "Skip ahead/behind rev-list walks (also when `status.aheadBehind=false`)")
	statusCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	statusCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	statusCmd.Flags().Bool("short", false, "One line per path: two-letter index/worktree code plus path (`??` / `!!` for untracked / ignored)")
	statusCmd.Flags().String("untracked", "", "How untracked paths are listed: `normal` (default) collapses wholly untracked directories; `all` lists every file; `no` omits untracked paths. When omitted, uses `status.showUntrackedFiles` from config")
	rootCmd.AddCommand(statusCmd)
}
