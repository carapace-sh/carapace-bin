package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var artifactRegistry_getTokenCmd = &cobra.Command{
	Use:   "get-token",
	Short: "Get a short-lived access token for the GitLab Artifact Registry. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifactRegistry_getTokenCmd).Standalone()

	artifactRegistry_getTokenCmd.Flags().String("duration", "15m0s", "How long the token should remain valid. Must be between 1s and 12h0m0s.")
	artifactRegistry_getTokenCmd.Flags().String("hostname", "", "GitLab hostname to request the token from. Defaults to the configured GitLab instance.")
	artifactRegistry_getTokenCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	artifactRegistry_getTokenCmd.Flags().StringP("output", "F", "text", "Format output as: text, json.")
	artifactRegistryCmd.AddCommand(artifactRegistry_getTokenCmd)

	carapace.Gen(artifactRegistry_getTokenCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"jq":       jq.ActionFilters(),
		"output":   carapace.ActionValues("text", "json"),
	})
}
