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
			"init", "absorb", "discard", "status", "tui", "stf", "uncommit", "amend",
			"squash", "move", "diff", "diff2", "edit", "show", "commit", "commit-empty",
			"push", "reword", "oplog-list", "oplog-snapshot", "restore", "undo", "redo",
			"gui", "open", "base-fetch", "base-check", "pull", "branch-new", "branch-delete",
			"branch-list", "branch-show", "branch-unapply", "branch-apply", "branch-update",
			"branch-move", "branch-tear-off",
		),
	})
}
