package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_inspectCmd = &cobra.Command{
	Use:   "inspect [ARTIFACT...]",
	Short: "Inspect an OCI artifact",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_inspectCmd).Standalone()

	artifactCmd.AddCommand(artifact_inspectCmd)
}
