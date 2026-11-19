package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_jobManifestPathsCmd = &cobra.Command{
	Use:   "jobManifestPaths",
	Short: "View skaffold jobManifestPath information defined in the specified skaffold configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_jobManifestPathsCmd).Standalone()

	applyCommonFlagCompletion(inspect_jobManifestPathsCmd)
	inspectCmd.AddCommand(inspect_jobManifestPathsCmd)
}
