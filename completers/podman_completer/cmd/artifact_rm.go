package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_rmCmd = &cobra.Command{
	Use:     "rm ARTIFACT",
	Short:   "Remove an OCI artifact",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_rmCmd).Standalone()

	artifactCmd.AddCommand(artifact_rmCmd)
}
