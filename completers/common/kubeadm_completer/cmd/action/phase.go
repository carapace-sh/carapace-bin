package action

import "github.com/carapace-sh/carapace"

func ActionPhases() carapace.Action {
	return carapace.ActionValuesDescribed(
		"addon", "Install required addons for passing conformance tests",
		"addon/coredns", "Install the CoreDNS addon to a Kubernetes cluster",
		"addon/kube-proxy", "Install the kube-proxy addon to a Kubernetes cluster",
		"bootstrap-token", "Generates bootstrap tokens used to join a node to a cluster",
		"certs", "Certificate generation",
		"certs/apiserver", "Generate the certificate for serving the Kubernetes API",
		"certs/apiserver-etcd-client", "Generate the certificate the apiserver uses to access etcd",
		"certs/apiserver-kubelet-client", "Generate the certificate for the API server to connect to kubelet",
		"certs/ca", "Generate the self-signed Kubernetes CA to provision identities for other Kubernetes components",
		"certs/etcd-ca", "Generate the self-signed CA to provision identities for etcd",
		"certs/etcd-healthcheck-client", "Generate the certificate for liveness probes to healthcheck etcd",
		"certs/etcd-peer", "Generate the certificate for etcd nodes to communicate with each other",
		"certs/etcd-server", "Generate the certificate for serving etcd",
		"certs/front-proxy-ca", "Generate the self-signed CA to provision identities for front proxy",
		"certs/front-proxy-client", "Generate the certificate for the front proxy client",
		"certs/sa", "Generate a private key for signing service account tokens along with its public key",
		"control-plane", "Generate all static Pod manifest files necessary to establish the control plane",
		"control-plane/apiserver", "Generates the kube-apiserver static Pod manifest",
		"control-plane/controller-manager", "Generates the kube-controller-manager static Pod manifest",
		"control-plane/scheduler", "Generates the kube-scheduler static Pod manifest",
		"etcd", "Generate static Pod manifest file for local etcd",
		"etcd/local", "Generate the static Pod manifest file for a local, single-node local etcd instance",
		"kubeconfig", "Generate all kubeconfig files necessary to establish the control plane and the admin kubeconfig file",
		"kubeconfig/admin", "Generate a kubeconfig file for the admin to use and for kubeadm itself",
		"kubeconfig/controller-manager", "Generate a kubeconfig file for the controller manager to use",
		"kubeconfig/kubelet", "Generate a kubeconfig file for the kubelet to use *only* for cluster bootstrapping purposes",
		"kubeconfig/scheduler", "Generate a kubeconfig file for the scheduler to use",
		"kubelet-finalize", "Updates settings relevant to the kubelet after TLS bootstrap",
		"kubelet-finalize/experimental-cert-rotation", "Enable kubelet client certificate rotation",
		"kubelet-start", "Write kubelet settings and (re)start the kubelet",
		"mark-control-plane", "Mark a node as a control-plane",
		"preflight", "Run pre-flight checks",
		"upload-certs", "Upload certificates to kubeadm-certs",
		"upload-config", "Upload the kubeadm and kubelet configuration to a ConfigMap",
		"upload-config/kubeadm", "Upload the kubeadm ClusterConfiguration to a ConfigMap",
		"upload-config/kubelet", "Upload the kubelet component config to a ConfigMap",
	)
}
