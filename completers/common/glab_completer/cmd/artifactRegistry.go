package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifactRegistryCmd = &cobra.Command{
	Use:     "artifact-registry <command>",
	Short:   "Authenticate with GitLab Artifact Registry. (EXPERIMENTAL)",
	Aliases: []string{"ar"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifactRegistryCmd).Standalone()

	rootCmd.AddCommand(artifactRegistryCmd)
}
