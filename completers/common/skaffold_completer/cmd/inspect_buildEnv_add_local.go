package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_buildEnv_add_localCmd = &cobra.Command{
	Use:   "local",
	Short: "Add a new 'local' build environment definition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_buildEnv_add_localCmd).Standalone()

	inspect_buildEnv_add_localCmd.Flags().String("concurrency", "-1", "number of artifacts to build concurrently. 0 means \"no-limit\"")
	inspect_buildEnv_add_localCmd.Flags().Bool("push", false, "Set to true to push images to a registry")
	inspect_buildEnv_add_localCmd.Flags().Bool("tryImportMissing", false, "Set to true to to attempt importing artifacts from Docker (either a local or remote registry) if not in the build cache")
	inspect_buildEnv_add_localCmd.Flags().Bool("useBuildkit", false, "Set to true to use BuildKit to build Docker images")
	inspect_buildEnv_add_localCmd.Flags().Bool("useDockerCLI", false, "Set to true to use 'docker' command-line interface instead of Docker Engine APIs")
	applyCommonFlagCompletion(inspect_buildEnv_add_localCmd)
	inspect_buildEnv_addCmd.AddCommand(inspect_buildEnv_add_localCmd)
}
