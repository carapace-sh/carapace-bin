package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var metricsCmd = &cobra.Command{
	Use:    "metrics",
	Short:  "INTERNAL: If metrics are permitted, this subcommand handles posthog event creation",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metricsCmd).Standalone()

	metricsCmd.Flags().String("command-name", "", "")
	metricsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	metricsCmd.Flags().String("props", "", "")
	metricsCmd.MarkFlagRequired("command-name")
	metricsCmd.MarkFlagRequired("props")
	rootCmd.AddCommand(metricsCmd)

	carapace.Gen(metricsCmd).FlagCompletion(carapace.ActionMap{
		"command-name": carapace.ActionValues(
			"absorb", "actions", "agent-log", "agent-setup", "alias-add", "alias-check", "alias-remove",
			"amend", "base-check", "base-fetch", "branch-apply", "branch-delete", "branch-list", "branch-move",
			"branch-new", "branch-show", "branch-tear-off", "branch-unapply", "branch-update", "clean", "comment",
			"commit", "commit-empty", "completions", "config", "diff", "disable-auto-merge", "discard", "edit",
			"enable-auto-merge", "external", "forge-auth", "forge-forget", "forge-list-users", "gui", "help",
			"init", "land", "mcp", "metrics", "move", "onboarding", "open", "oplog-list", "oplog-snapshot",
			"pick", "pr-new", "pr-template", "pull", "push", "redo", "refresh-remote-data", "resolve", "restore",
			"reword", "set-review-draft", "set-review-ready", "setup", "show", "skill", "skill-check",
			"skill-install", "split", "squash", "status", "stf", "switch", "teardown", "tui", "uncommit",
			"undo", "update", "update-check", "update-install", "update-suppress", "worktree-archive",
			"worktree-list", "worktree-new", "worktree-remove", "worktree-unarchive",
		),
	})
}
