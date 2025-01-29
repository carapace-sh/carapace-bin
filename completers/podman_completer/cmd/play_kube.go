package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var play_kubeCmd = &cobra.Command{
	Use:    "kube [options] KUBEFILE|-",
	Short:  "Play a pod or volume based on Kubernetes YAML",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(play_kubeCmd).Standalone()

	play_kubeCmd.Flags().StringSlice("annotation", []string{}, "Add annotations to pods (key=value)")
	play_kubeCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	play_kubeCmd.Flags().Bool("build", false, "Build all images in a YAML (given Containerfiles exist)")
	play_kubeCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	play_kubeCmd.Flags().StringSlice("configmap", []string{}, "`Pathname` of a YAML file containing a kubernetes configmap")
	play_kubeCmd.Flags().String("context-dir", "", "Path to top level of context directory")
	play_kubeCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	play_kubeCmd.Flags().Bool("down", false, "Stop pods defined in the YAML file")
	play_kubeCmd.Flags().Bool("force", false, "Remove volumes as part of --down")
	play_kubeCmd.Flags().StringSlice("ip", []string{}, "Static IP addresses to assign to the pods")
	play_kubeCmd.Flags().String("log-driver", "", "Logging driver for the container")
	play_kubeCmd.Flags().StringSlice("log-opt", []string{}, "Logging driver options")
	play_kubeCmd.Flags().StringSlice("mac-address", []string{}, "Static MAC addresses to assign to the pods")
	play_kubeCmd.Flags().StringSlice("network", []string{}, "Connect pod to network(s) or network mode")
	play_kubeCmd.Flags().Bool("no-hostname", false, "Do not create /etc/hostname within the container, instead use the version from the image")
	play_kubeCmd.Flags().Bool("no-hosts", false, "Do not create /etc/hosts within the pod's containers, instead use the version from the image")
	play_kubeCmd.Flags().Bool("no-trunc", false, "Use annotations that are not truncated to the Kubernetes maximum length of 63 characters")
	play_kubeCmd.Flags().StringSlice("publish", []string{}, "Publish a container's port, or a range of ports, to the host")
	play_kubeCmd.Flags().Bool("publish-all", false, "Whether to publish all ports defined in the K8S YAML file (containerPort, hostPort), if false only hostPort will be published")
	play_kubeCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when pulling images")
	play_kubeCmd.Flags().Bool("replace", false, "Delete and recreate pods defined in the YAML file")
	play_kubeCmd.Flags().String("seccomp-profile-root", "", "Directory path for seccomp profiles")
	play_kubeCmd.Flags().Bool("service-container", false, "Starts a service container before all pods")
	play_kubeCmd.Flags().String("service-exit-code-propagation", "", "Exit-code propagation of the service container")
	play_kubeCmd.Flags().String("signature-policy", "", "`Pathname` of signature policy file (not usually used)")
	play_kubeCmd.Flags().Bool("start", false, "Start the pod after creating it")
	play_kubeCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	play_kubeCmd.Flags().String("userns", "", "User namespace to use")
	play_kubeCmd.Flags().BoolP("wait", "w", false, "Clean up all objects created when a SIGTERM is received or pods exit")
	play_kubeCmd.Flag("down").Hidden = true
	play_kubeCmd.Flag("no-trunc").Hidden = true
	play_kubeCmd.Flag("service-container").Hidden = true
	play_kubeCmd.Flag("service-exit-code-propagation").Hidden = true
	play_kubeCmd.Flag("signature-policy").Hidden = true
	playCmd.AddCommand(play_kubeCmd)
}
