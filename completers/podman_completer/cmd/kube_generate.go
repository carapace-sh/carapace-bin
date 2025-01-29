package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kube_generateCmd = &cobra.Command{
	Use:   "generate [options] {CONTAINER...|POD...|VOLUME...}",
	Short: "Generate Kubernetes YAML from containers, pods or volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_generateCmd).Standalone()

	kube_generateCmd.Flags().StringP("filename", "f", "", "Write output to the specified path")
	kube_generateCmd.Flags().Bool("no-trunc", false, "Don't truncate annotations to Kubernetes length (63 chars)")
	kube_generateCmd.Flags().Bool("podman-only", false, "Add podman-only reserved annotations to the generated YAML file (Cannot be used by Kubernetes)")
	kube_generateCmd.Flags().StringP("replicas", "r", "", "Set the replicas number for Deployment kind")
	kube_generateCmd.Flags().BoolP("service", "s", false, "Generate YAML for a Kubernetes service object")
	kube_generateCmd.Flags().StringP("type", "t", "", "Generate YAML for the given Kubernetes kind")
	kube_generateCmd.Flag("no-trunc").Hidden = true
	kubeCmd.AddCommand(kube_generateCmd)
}
