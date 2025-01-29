package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_addCmd = &cobra.Command{
	Use:   "add ARTIFACT PATH [...PATH]",
	Short: "Add an OCI artifact to the local store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_addCmd).Standalone()

	artifactCmd.AddCommand(artifact_addCmd)
}
