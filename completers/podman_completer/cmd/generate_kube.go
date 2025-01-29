package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_kubeCmd = &cobra.Command{
	Use:   "kube [options] {CONTAINER...|POD...|VOLUME...}",
	Short: "Generate Kubernetes YAML from containers, pods or volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_kubeCmd).Standalone()

	generate_kubeCmd.Flags().StringP("filename", "f", "", "Write output to the specified path")
	generate_kubeCmd.Flags().Bool("no-trunc", false, "Don't truncate annotations to Kubernetes length (63 chars)")
	generate_kubeCmd.Flags().Bool("podman-only", false, "Add podman-only reserved annotations to the generated YAML file (Cannot be used by Kubernetes)")
	generate_kubeCmd.Flags().StringP("replicas", "r", "", "Set the replicas number for Deployment kind")
	generate_kubeCmd.Flags().BoolP("service", "s", false, "Generate YAML for a Kubernetes service object")
	generate_kubeCmd.Flags().StringP("type", "t", "", "Generate YAML for the given Kubernetes kind")
	generate_kubeCmd.Flag("no-trunc").Hidden = true
	generateCmd.AddCommand(generate_kubeCmd)
}
