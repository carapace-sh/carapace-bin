package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var version_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply all the deferred version bumps at once",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(version_applyCmd).Standalone()

	version_applyCmd.Flags().Bool("all", false, "Apply the deferred version changes on all workspaces")
	version_applyCmd.Flags().Bool("dry-run", false, "Print the versions without actually generating the package archive")
	version_applyCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	version_applyCmd.Flags().Bool("prerelease", false, "Add a prerelease identifier to new versions")
	version_applyCmd.Flags().BoolP("recursive", "R", false, "Release the transitive workspaces as well")
	versionCmd.AddCommand(version_applyCmd)
}
