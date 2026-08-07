package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orbit_remote_toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Show the GitLab Knowledge Graph MCP tool manifest. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_remote_toolsCmd).Standalone()

	orbit_remote_toolsCmd.Flags().String("hostname", "", "GitLab hostname to query. Defaults to the current repository's host or `gitlab.com`.")
	orbit_remote_toolsCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	orbit_remoteCmd.AddCommand(orbit_remote_toolsCmd)
}
