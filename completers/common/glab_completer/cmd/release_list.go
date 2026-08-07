package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List releases in a repository.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_listCmd).Standalone()

	release_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	release_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	release_listCmd.Flags().StringP("page", "p", "", "Page number.")
	release_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	release_listCmd.Flags().StringP("tag", "t", "", "Filter releases by tag <name>.")
	release_listCmd.Flag("tag").Hidden = true
	releaseCmd.AddCommand(release_listCmd)

	carapace.Gen(release_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
		"tag":    action.ActionReleases(release_listCmd),
	})
}
