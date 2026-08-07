package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orbit_remoteCmd = &cobra.Command{
	Use:     "remote <command> [flags]",
	Short:   "Interact with the remote GitLab Knowledge Graph. (EXPERIMENTAL)",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_remoteCmd).Standalone()

	orbitCmd.AddCommand(orbit_remoteCmd)
}
