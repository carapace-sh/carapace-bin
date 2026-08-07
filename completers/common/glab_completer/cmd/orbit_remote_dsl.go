package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orbit_remote_dslCmd = &cobra.Command{
	Use:   "dsl",
	Short: "Show the GitLab Knowledge Graph query DSL JSON Schema. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_remote_dslCmd).Standalone()

	orbit_remote_dslCmd.Flags().String("hostname", "", "GitLab hostname to query. Defaults to the current repository's host or `gitlab.com`.")
	orbit_remoteCmd.AddCommand(orbit_remote_dslCmd)
}
