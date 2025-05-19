package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kubeadm_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join [api-server-endpoint]",
	Short: "Run this on any machine you wish to join an existing cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinCmd).Standalone()

	joinCmd.Flags().String("apiserver-advertise-address", "", "If the node should host a new control plane instance, the IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	joinCmd.Flags().String("apiserver-bind-port", "", "If the node should host a new control plane instance, the port for the API Server to bind to.")
	joinCmd.Flags().String("certificate-key", "", "Use this key to decrypt the certificate secrets uploaded by init. The certificate key is a hex encoded string that is an AES key of size 32 bytes.")
	joinCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	joinCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	joinCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	joinCmd.Flags().String("discovery-file", "", "For file-based discovery, a file or URL from which to load cluster information.")
	joinCmd.Flags().String("discovery-token", "", "For token-based discovery, the token used to validate cluster information fetched from the API server.")
	joinCmd.Flags().StringSlice("discovery-token-ca-cert-hash", nil, "For token-based discovery, validate that the root CA public key matches this hash (format: \"<type>:<value>\").")
	joinCmd.Flags().Bool("discovery-token-unsafe-skip-ca-verification", false, "For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.")
	joinCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	joinCmd.Flags().StringSlice("ignore-preflight-errors", nil, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	joinCmd.Flags().String("node-name", "", "Specify the node name.")
	joinCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	joinCmd.Flags().StringSlice("skip-phases", nil, "List of phases to be skipped")
	joinCmd.Flags().String("tls-bootstrap-token", "", "Specify the token used to temporarily authenticate with the Kubernetes Control Plane while joining the node.")
	joinCmd.Flags().String("token", "", "Use this token for both discovery-token and tls-bootstrap-token when those values are not provided.")
	rootCmd.AddCommand(joinCmd)

	carapace.Gen(joinCmd).FlagCompletion(carapace.ActionMap{
		"apiserver-bind-port":     net.ActionPorts(),
		"config":                  carapace.ActionFiles(),
		"cri-socket":              carapace.ActionFiles(),
		"discovery-file":          carapace.ActionFiles(),
		"ignore-preflight-errors": action.ActionChecks().UniqueList(","),
		"patches":                 carapace.ActionDirectories(),
		"skip-phases":             action.ActionPhases(),
	})
}
