package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var join_phase_controlPlanePrepare_certsCmd = &cobra.Command{
	Use:   "certs [api-server-endpoint]",
	Short: "Generate the certificates for the new control plane components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_controlPlanePrepare_certsCmd).Standalone()

	join_phase_controlPlanePrepare_certsCmd.Flags().String("apiserver-advertise-address", "", "If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	join_phase_controlPlanePrepare_certsCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	join_phase_controlPlanePrepare_certsCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	join_phase_controlPlanePrepare_certsCmd.Flags().String("discovery-file", "", "For file-based discovery, a file or URL from which to load cluster information.")
	join_phase_controlPlanePrepare_certsCmd.Flags().String("discovery-token", "", "For token-based discovery, the token used to validate cluster information fetched from the API server.")
	join_phase_controlPlanePrepare_certsCmd.Flags().StringSlice("discovery-token-ca-cert-hash", nil, "For token-based discovery, validate that the root CA public key matches this hash (format: \"<type>:<value>\").")
	join_phase_controlPlanePrepare_certsCmd.Flags().Bool("discovery-token-unsafe-skip-ca-verification", false, "For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.")
	join_phase_controlPlanePrepare_certsCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	join_phase_controlPlanePrepare_certsCmd.Flags().String("node-name", "", "Specify the node name.")
	join_phase_controlPlanePrepare_certsCmd.Flags().String("tls-bootstrap-token", "", "Specify the token used to temporarily authenticate with the Kubernetes Control Plane while joining the node.")
	join_phase_controlPlanePrepare_certsCmd.Flags().String("token", "", "Use this token for both discovery-token and tls-bootstrap-token when those values are not provided.")
	join_phase_controlPlanePrepareCmd.AddCommand(join_phase_controlPlanePrepare_certsCmd)

	carapace.Gen(join_phase_controlPlanePrepare_certsCmd).FlagCompletion(carapace.ActionMap{
		"config":         carapace.ActionFiles(),
		"discovery-file": carapace.ActionFiles(),
	})
}
