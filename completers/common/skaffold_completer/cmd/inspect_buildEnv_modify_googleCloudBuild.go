package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_buildEnv_modify_googleCloudBuildCmd = &cobra.Command{
	Use:   "googleCloudBuild",
	Short: "Modify an existing GoogleCloudBuild build environment definition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_buildEnv_modify_googleCloudBuildCmd).Standalone()

	inspect_buildEnv_modify_googleCloudBuildCmd.Flags().String("concurrency", "-1", "number of artifacts to build concurrently. 0 means \"no-limit\"")
	inspect_buildEnv_modify_googleCloudBuildCmd.Flags().String("diskSizeGb", "", "Disk size of the VM that runs the build")
	inspect_buildEnv_modify_googleCloudBuildCmd.Flags().String("logStreamingOption", "", "Specifies the log streaming specifies behavior when writing build logs to Google Cloud Storage for GCB")
	inspect_buildEnv_modify_googleCloudBuildCmd.Flags().String("logging", "", "Specifies the logging mode for GCB")
	inspect_buildEnv_modify_googleCloudBuildCmd.Flags().String("machineType", "", "Type of VM that runs the build")
	inspect_buildEnv_modify_googleCloudBuildCmd.Flags().String("projectId", "", "ID of the Cloud Platform Project.")
	inspect_buildEnv_modify_googleCloudBuildCmd.Flags().String("timeout", "", "Build timeout (in seconds)")
	inspect_buildEnv_modify_googleCloudBuildCmd.Flags().String("workerPool", "", "Configures a pool of workers to run the build")
	applyCommonFlagCompletion(inspect_buildEnv_modify_googleCloudBuildCmd)
	inspect_buildEnv_modifyCmd.AddCommand(inspect_buildEnv_modify_googleCloudBuildCmd)
}
