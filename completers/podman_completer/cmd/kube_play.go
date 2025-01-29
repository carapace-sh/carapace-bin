package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kube_playCmd = &cobra.Command{
	Use:   "play [options] KUBEFILE|-",
	Short: "Play a pod or volume based on Kubernetes YAML",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_playCmd).Standalone()

	kube_playCmd.Flags().StringSlice("annotation", []string{}, "Add annotations to pods (key=value)")
	kube_playCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	kube_playCmd.Flags().Bool("build", false, "Build all images in a YAML (given Containerfiles exist)")
	kube_playCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	kube_playCmd.Flags().StringSlice("configmap", []string{}, "`Pathname` of a YAML file containing a kubernetes configmap")
	kube_playCmd.Flags().String("context-dir", "", "Path to top level of context directory")
	kube_playCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	kube_playCmd.Flags().Bool("down", false, "Stop pods defined in the YAML file")
	kube_playCmd.Flags().Bool("force", false, "Remove volumes as part of --down")
	kube_playCmd.Flags().StringSlice("ip", []string{}, "Static IP addresses to assign to the pods")
	kube_playCmd.Flags().String("log-driver", "", "Logging driver for the container")
	kube_playCmd.Flags().StringSlice("log-opt", []string{}, "Logging driver options")
	kube_playCmd.Flags().StringSlice("mac-address", []string{}, "Static MAC addresses to assign to the pods")
	kube_playCmd.Flags().StringSlice("network", []string{}, "Connect pod to network(s) or network mode")
	kube_playCmd.Flags().Bool("no-hostname", false, "Do not create /etc/hostname within the container, instead use the version from the image")
	kube_playCmd.Flags().Bool("no-hosts", false, "Do not create /etc/hosts within the pod's containers, instead use the version from the image")
	kube_playCmd.Flags().Bool("no-trunc", false, "Use annotations that are not truncated to the Kubernetes maximum length of 63 characters")
	kube_playCmd.Flags().StringSlice("publish", []string{}, "Publish a container's port, or a range of ports, to the host")
	kube_playCmd.Flags().Bool("publish-all", false, "Whether to publish all ports defined in the K8S YAML file (containerPort, hostPort), if false only hostPort will be published")
	kube_playCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pulling images")
	kube_playCmd.Flags().Bool("replace", false, "Delete and recreate pods defined in the YAML file")
	kube_playCmd.Flags().String("seccomp-profile-root", "", "Directory path for seccomp profiles")
	kube_playCmd.Flags().Bool("service-container", false, "Starts a service container before all pods")
	kube_playCmd.Flags().String("service-exit-code-propagation", "", "Exit-code propagation of the service container")
	kube_playCmd.Flags().String("signature-policy", "", "`Pathname` of signature policy file (not usually used)")
	kube_playCmd.Flags().Bool("start", false, "Start the pod after creating it")
	kube_playCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	kube_playCmd.Flags().String("userns", "", "User namespace to use")
	kube_playCmd.Flags().BoolP("wait", "w", false, "Clean up all objects created when a SIGTERM is received or pods exit")
	kube_playCmd.Flag("down").Hidden = true
	kube_playCmd.Flag("no-trunc").Hidden = true
	kube_playCmd.Flag("service-container").Hidden = true
	kube_playCmd.Flag("service-exit-code-propagation").Hidden = true
	kube_playCmd.Flag("signature-policy").Hidden = true
	kubeCmd.AddCommand(kube_playCmd)
}
