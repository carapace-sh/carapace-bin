package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_publish_catalogCmd = &cobra.Command{
	Use:   "catalog <tag-name>",
	Short: "[EXPERIMENTAL] Publishes CI/CD components to the catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_publish_catalogCmd).Standalone()

	repo_publishCmd.AddCommand(repo_publish_catalogCmd)

	carapace.Gen(repo_publishCmd).PositionalCompletion(
		action.ActionTags(repo_publishCmd),
	)
}
