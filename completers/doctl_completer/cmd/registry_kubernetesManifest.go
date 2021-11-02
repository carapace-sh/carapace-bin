package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_kubernetesManifestCmd = &cobra.Command{
	Use:   "kubernetes-manifest",
	Short: "Generate a Kubernetes secret manifest for a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_kubernetesManifestCmd).Standalone()
	registry_kubernetesManifestCmd.Flags().String("name", "", "The secret name to create. Defaults to the registry name prefixed with \"registry-\"")
	registry_kubernetesManifestCmd.Flags().String("namespace", "kube-system", "The Kubernetes namespace to hold the secret")
	registryCmd.AddCommand(registry_kubernetesManifestCmd)
}
