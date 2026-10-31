package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_buildEnv_add_clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Add a new 'cluster' build environment definition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_buildEnv_add_clusterCmd).Standalone()

	inspect_buildEnv_add_clusterCmd.Flags().String("concurrency", "-1", "number of artifacts to build concurrently. 0 means \"no-limit\"")
	inspect_buildEnv_add_clusterCmd.Flags().String("dockerConfigPath", "", "Path to the docker config.json.")
	inspect_buildEnv_add_clusterCmd.Flags().String("dockerConfigSecretName", "", "Kubernetes secret that contains the config.json Docker configuration.")
	inspect_buildEnv_add_clusterCmd.Flags().String("namespace", "", "Kubernetes namespace.")
	inspect_buildEnv_add_clusterCmd.Flags().String("pullSecretMountPath", "", "Path the pull secret will be mounted at within the running container.")
	inspect_buildEnv_add_clusterCmd.Flags().String("pullSecretName", "", "Name of the Kubernetes secret for pulling base images and pushing the final image.")
	inspect_buildEnv_add_clusterCmd.Flags().String("pullSecretPath", "", "Path to the Google Cloud service account secret key file.")
	inspect_buildEnv_add_clusterCmd.Flags().Bool("randomDockerConfigSecret", false, "Adds a random UUID postfix to the default name of the docker secret to facilitate parallel builds.")
	inspect_buildEnv_add_clusterCmd.Flags().Bool("randomPullSecret", false, "Adds a random UUID postfix to the default name of the pull secret to facilitate parallel builds.")
	inspect_buildEnv_add_clusterCmd.Flags().String("runAsUser", "-1", "Defines the UID to request for running the container.")
	inspect_buildEnv_add_clusterCmd.Flags().String("serviceAccount", "", "Kubernetes service account to use for the pod.")
	inspect_buildEnv_add_clusterCmd.Flags().String("timeout", "", "Build timeout (in seconds)")
	applyCommonFlagCompletion(inspect_buildEnv_add_clusterCmd)
	inspect_buildEnv_addCmd.AddCommand(inspect_buildEnv_add_clusterCmd)
}
