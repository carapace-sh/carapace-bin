package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var orbit_remote_schemaCmd = &cobra.Command{
	Use:   "schema [node...]",
	Short: "Show the GitLab Knowledge Graph ontology. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_remote_schemaCmd).Standalone()

	orbit_remote_schemaCmd.Flags().String("hostname", "", "GitLab hostname to query. Defaults to the current repository's host or `gitlab.com`.")
	orbit_remote_schemaCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	orbit_remoteCmd.AddCommand(orbit_remote_schemaCmd)

	carapace.Gen(orbit_remote_schemaCmd).FlagCompletion(carapace.ActionMap{
		"jq": jq.ActionFilters(),
	})
}
