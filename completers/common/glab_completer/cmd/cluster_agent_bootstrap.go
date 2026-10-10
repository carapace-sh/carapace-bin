package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cluster_agent_bootstrapCmd = &cobra.Command{
	Use:     "bootstrap <agent-name> [flags]",
	Short:   "Bootstrap a GitLab Agent for Kubernetes in a project.",
	Aliases: []string{"bs"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_bootstrapCmd).Standalone()

	cluster_agent_bootstrapCmd.Flags().String("commit-author-email", "noreply@glab.gitlab.com", "The Git commit author email to use. Conflicts with '--use-api-commit-author'.")
	cluster_agent_bootstrapCmd.Flags().String("commit-author-name", "glab", "The Git commit author name to use. Conflicts with '--use-api-commit-author'.")
	cluster_agent_bootstrapCmd.Flags().Bool("create-environment", true, "Create an environment for the GitLab Agent.")
	cluster_agent_bootstrapCmd.Flags().Bool("create-flux-environment", true, "Create an environment for FluxCD. Affects only the environment creation, not the use of Flux itself. Flux is always required for the bootstrap process.")
	cluster_agent_bootstrapCmd.Flags().String("environment-flux-resource-path", "helm.toolkit.fluxcd.io/v2beta1/namespaces/<helm-release-namespace>/helmreleases/<helm-release-name>", "Flux resource path of the environment for the GitLab Agent.")
	cluster_agent_bootstrapCmd.Flags().String("environment-name", "<helm-release-namespace>/<helm-release-name>", "Name of the environment for the GitLab Agent.")
	cluster_agent_bootstrapCmd.Flags().String("environment-namespace", "<helm-release-namespace>", "Kubernetes namespace of the environment for the GitLab Agent.")
	cluster_agent_bootstrapCmd.Flags().String("flux-environment-flux-resource-path", "kustomize.toolkit.fluxcd.io/v1/namespaces/flux-system/kustomizations/flux-system", "Flux resource path of the environment for FluxCD.")
	cluster_agent_bootstrapCmd.Flags().String("flux-environment-name", "<flux-source-namespace>/<flux-source-name>", "Name of the environment for FluxCD.")
	cluster_agent_bootstrapCmd.Flags().String("flux-environment-namespace", "<flux-source-namespace>", "Kubernetes namespace of the environment for FluxCD.")
	cluster_agent_bootstrapCmd.Flags().String("flux-source-name", "flux-system", "Flux source name.")
	cluster_agent_bootstrapCmd.Flags().String("flux-source-namespace", "flux-system", "Flux source namespace.")
	cluster_agent_bootstrapCmd.Flags().String("flux-source-type", "git", "Source type of the flux-system, like Git, OCI, or Helm.")
	cluster_agent_bootstrapCmd.Flags().String("gitlab-agent-token-secret-name", "gitlab-agent-token", "Name of the Secret where the token for the GitLab Agent is stored. The helm-release-target-namespace is implied for the namespace of the Secret.")
	cluster_agent_bootstrapCmd.Flags().String("helm-release-filepath", "gitlab-agent-helm-release.yaml", "File path within the GitLab Agent project to commit the Flux HelmRelease to.")
	cluster_agent_bootstrapCmd.Flags().String("helm-release-name", "gitlab-agent", "Name of the Flux HelmRelease manifest.")
	cluster_agent_bootstrapCmd.Flags().String("helm-release-namespace", "flux-system", "Namespace of the Flux HelmRelease manifest.")
	cluster_agent_bootstrapCmd.Flags().String("helm-release-target-namespace", "gitlab-agent", "Namespace of the GitLab Agent deployment.")
	cluster_agent_bootstrapCmd.Flags().StringSlice("helm-release-values", nil, "Local path to values.yaml files. Multiple files can be comma-separated or specified by repeating the flag.")
	cluster_agent_bootstrapCmd.Flags().StringSlice("helm-release-values-from", nil, "Kubernetes object reference that contains the values.yaml data key in the format '<kind>/<name>', where 'kind' must be one of: (Secret, ConfigMap). Multiple references can be comma-separated or specified by repeating the flag.")
	cluster_agent_bootstrapCmd.Flags().String("helm-repository-address", "https://charts.gitlab.io", "Address of the HelmRepository.")
	cluster_agent_bootstrapCmd.Flags().String("helm-repository-filepath", "gitlab-helm-repository.yaml", "File path within the GitLab Agent project to commit the Flux HelmRepository to.")
	cluster_agent_bootstrapCmd.Flags().String("helm-repository-name", "gitlab", "Name of the Flux HelmRepository manifest.")
	cluster_agent_bootstrapCmd.Flags().String("helm-repository-namespace", "flux-system", "Namespace of the Flux HelmRepository manifest.")
	cluster_agent_bootstrapCmd.Flags().StringP("manifest-branch", "b", "", "Branch to commit the Flux Manifests to. Defaults to the project default branch.")
	cluster_agent_bootstrapCmd.Flags().StringP("manifest-path", "p", "", "Location of directory in Git repository for storing the GitLab Agent for Kubernetes Helm resources.")
	cluster_agent_bootstrapCmd.Flags().Bool("no-reconcile", false, "Do not trigger Flux reconciliation for GitLab Agent for Kubernetes Flux resource.")
	cluster_agent_bootstrapCmd.Flags().Bool("use-api-commit-author", false, "When creating Git commits use the user from the authenticated API request. Conflicts with '--commit-author-name' and '--commit-author-email'.")
	cluster_agentCmd.AddCommand(cluster_agent_bootstrapCmd)

	carapace.Gen(cluster_agent_bootstrapCmd).FlagCompletion(carapace.ActionMap{
		"flux-source-type":    carapace.ActionValues("Git", "OCI", "Helm"),
		"helm-release-values": carapace.ActionFiles(),
		"manifest-branch":     action.ActionBranches(cluster_agent_bootstrapCmd),
	})
}
