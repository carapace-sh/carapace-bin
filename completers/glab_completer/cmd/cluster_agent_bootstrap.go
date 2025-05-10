package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_bootstrapCmd = &cobra.Command{
	Use:     "bootstrap agent-name [flags]",
	Short:   "Bootstrap a GitLab Agent for Kubernetes in a project.",
	Aliases: []string{"bs"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_bootstrapCmd).Standalone()

	cluster_agent_bootstrapCmd.Flags().Bool("create-environment", false, "Create an Environment for the GitLab Agent.")
	cluster_agent_bootstrapCmd.Flags().String("environment-flux-resource-path", "", "Flux Resource Path of the Environment for the GitLab Agent.")
	cluster_agent_bootstrapCmd.Flags().String("environment-name", "", "Name of the Environment for the GitLab Agent.")
	cluster_agent_bootstrapCmd.Flags().String("environment-namespace", "", "Kubernetes namespace of the Environment for the GitLab Agent.")
	cluster_agent_bootstrapCmd.Flags().String("flux-source-name", "", "Flux source name.")
	cluster_agent_bootstrapCmd.Flags().String("flux-source-namespace", "", "Flux source namespace.")
	cluster_agent_bootstrapCmd.Flags().String("flux-source-type", "", "Source type of the flux-system, e.g. git, oci, helm, ...")
	cluster_agent_bootstrapCmd.Flags().String("gitlab-agent-token-secret-name", "", "Name of the Secret where the token for the GitLab Agent is stored. The helm-release-target-namespace is implied for the namespace of the Secret.")
	cluster_agent_bootstrapCmd.Flags().String("helm-release-filepath", "", "Filepath within the GitLab Agent project to commit the Flux HelmRelease to.")
	cluster_agent_bootstrapCmd.Flags().String("helm-release-name", "", "Name of the Flux HelmRelease manifest.")
	cluster_agent_bootstrapCmd.Flags().String("helm-release-namespace", "", "Namespace of the Flux HelmRelease manifest.")
	cluster_agent_bootstrapCmd.Flags().String("helm-release-target-namespace", "", "Namespace of the GitLab Agent deployment.")
	cluster_agent_bootstrapCmd.Flags().StringSlice("helm-release-values", nil, "Local path to values.yaml files")
	cluster_agent_bootstrapCmd.Flags().StringSlice("helm-release-values-from", nil, "Kubernetes object reference that contains the values.yaml data key in the format '<kind>/<name>', where kind must be one of: (Secret,ConfigMap)")
	cluster_agent_bootstrapCmd.Flags().String("helm-repository-filepath", "", "Filepath within the GitLab Agent project to commit the Flux HelmRepository to.")
	cluster_agent_bootstrapCmd.Flags().String("helm-repository-name", "", "Name of the Flux HelmRepository manifest.")
	cluster_agent_bootstrapCmd.Flags().String("helm-repository-namespace", "", "Namespace of the Flux HelmRepository manifest.")
	cluster_agent_bootstrapCmd.Flags().StringP("manifest-branch", "b", "", "Branch to commit the Flux Manifests to. (default to the project default branch)")
	cluster_agent_bootstrapCmd.Flags().StringP("manifest-path", "p", "", "Location of directory in Git repository for storing the GitLab Agent for Kubernetes Helm resources.")
	cluster_agent_bootstrapCmd.Flags().Bool("no-reconcile", false, "Do not trigger Flux reconciliation for GitLab Agent for Kubernetes Flux resource.")
	cluster_agentCmd.AddCommand(cluster_agent_bootstrapCmd)

	// TODO flag completion
}
