package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var orbit_remote_graphStatusCmd = &cobra.Command{
	Use:   "graph-status",
	Short: "Show indexing progress for a namespace or project. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_remote_graphStatusCmd).Standalone()

	orbit_remote_graphStatusCmd.Flags().StringP("format", "f", "", "Response format: `raw` (structured JSON) or `llm` (compact, intended for agents).")
	orbit_remote_graphStatusCmd.Flags().String("full-path", "", "Full path of a project or group, such as `gitlab-org/gitlab`. Cannot be used with the ID flags.")
	orbit_remote_graphStatusCmd.Flags().String("hostname", "", "GitLab hostname to query. Defaults to the current repository's host or `gitlab.com`.")
	orbit_remote_graphStatusCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	orbit_remote_graphStatusCmd.Flags().String("namespace-id", "", "Namespace (group) ID to inspect. Cannot be used with --project-id or --full-path.")
	orbit_remote_graphStatusCmd.Flags().String("project-id", "", "Project ID to inspect. Cannot be used with --namespace-id or --full-path.")
	orbit_remote_graphStatusCmd.Flags().String("response-format", "", "Server response format: `raw` (structured JSON) or `llm` (compact GOON/TOON for agents).")
	orbit_remote_graphStatusCmd.Flag("format").Hidden = true
	orbit_remoteCmd.AddCommand(orbit_remote_graphStatusCmd)

	carapace.Gen(orbit_remote_graphStatusCmd).FlagCompletion(carapace.ActionMap{
		"format":          carapace.ActionValues("raw", "llm"),
		"hostname":        action.ActionConfigHosts(),
		"response-format": carapace.ActionValues("raw", "llm"),
	})
}
