package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orbit_remote_queryCmd = &cobra.Command{
	Use:   "query [file|-]",
	Short: "Execute a GitLab Knowledge Graph query. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_remote_queryCmd).Standalone()

	orbit_remote_queryCmd.Flags().StringP("format", "f", "", "Response format: `llm` (compact, intended for agents) or `raw` (structured JSON).")
	orbit_remote_queryCmd.Flags().String("hostname", "", "GitLab hostname to query. Defaults to the current repository's host or `gitlab.com`.")
	orbit_remote_queryCmd.Flags().String("response-format", "", "Server response format: `llm` (compact GOON/TOON for agents) or `raw` (structured JSON).")
	orbit_remote_queryCmd.Flag("format").Hidden = true
	orbit_remoteCmd.AddCommand(orbit_remote_queryCmd)
}
