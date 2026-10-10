package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var artifactRegistry_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check your access to the GitLab Artifact Registry. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifactRegistry_statusCmd).Standalone()

	artifactRegistry_statusCmd.Flags().String("hostname", "", "GitLab hostname to check. Defaults to the configured GitLab instance.")
	artifactRegistry_statusCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	artifactRegistry_statusCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	artifactRegistryCmd.AddCommand(artifactRegistry_statusCmd)

	carapace.Gen(artifactRegistry_statusCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"jq":       jq.ActionFilters(),
		"output":   carapace.ActionValues("text", "json"),
	})
}
