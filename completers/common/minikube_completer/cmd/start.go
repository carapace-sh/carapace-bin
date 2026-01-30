package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Starts a local Kubernetes cluster",
	GroupID: "basic",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	startCmd.Flags().StringSlice("addons", nil, "Enable one or more addons, in a comma-separated format. See `minikube addons list` for a list of valid addon names.")
	startCmd.Flags().StringSlice("apiserver-ips", nil, "A set of apiserver IP Addresses which are used in the generated certificate for kubernetes.  This can be used if you want to make the apiserver available from outside the machine")
	startCmd.Flags().String("apiserver-name", "", "The authoritative apiserver hostname for apiserver certificates and connectivity. This can be used if you want to make the apiserver available from outside the machine")
	startCmd.Flags().StringSlice("apiserver-names", nil, "A set of apiserver names which are used in the generated certificate for kubernetes.  This can be used if you want to make the apiserver available from outside the machine")
	startCmd.Flags().String("apiserver-port", "", "The apiserver listening port")
	startCmd.Flags().String("auto-pause-interval", "", "Duration of inactivity before the minikube VM is paused (default 1m0s)")
	startCmd.Flags().Bool("auto-update-drivers", false, "If set, automatically updates drivers to the latest version. Defaults to true.")
	startCmd.Flags().String("base-image", "", "The base image to use for docker/podman drivers. Intended for local development.")
	startCmd.Flags().String("binary-mirror", "", "Location to fetch kubectl, kubelet, & kubeadm binaries from.")
	startCmd.Flags().Bool("cache-images", false, "If true, cache docker images for the current bootstrapper and load them into the machine. Always false with --driver=none.")
	startCmd.Flags().String("cert-expiration", "", "Duration until minikube certificate expiration, defaults to three years (26280h).")
	startCmd.Flags().String("cni", "", "CNI plug-in to use. Valid options: auto, bridge, calico, cilium, flannel, kindnet, or path to a CNI manifest (default: auto)")
	startCmd.Flags().StringP("container-runtime", "c", "", "The container runtime to be used. Valid options: docker, cri-o, containerd (default: auto)")
	startCmd.Flags().String("cpus", "", "Number of CPUs allocated to Kubernetes. Use \"max\" to use the maximum number of CPUs. Use \"no-limit\" to not specify a limit (Docker/Podman only)")
	startCmd.Flags().String("cri-socket", "", "The cri socket path to be used.")
	startCmd.Flags().Bool("delete-on-failure", false, "If set, delete the current cluster if start fails and try again. Defaults to false.")
	startCmd.Flags().Bool("disable-coredns-log", false, "If set, disable CoreDNS verbose logging. Defaults to false.")
	startCmd.Flags().Bool("disable-driver-mounts", false, "Disables the filesystem mounts provided by the hypervisors")
	startCmd.Flags().Bool("disable-metrics", false, "If set, disables metrics reporting (CPU and memory usage), this can improve CPU usage. Defaults to false.")
	startCmd.Flags().Bool("disable-optimizations", false, "If set, disables optimizations that are set for local Kubernetes. Including decreasing CoreDNS replicas from 2 to 1. Defaults to false.")
	startCmd.Flags().String("disk-size", "", "Disk size allocated to the minikube VM (format: <number>[<unit>], where unit = b, k, m or g).")
	startCmd.Flags().String("dns-domain", "", "The cluster dns domain name used in the Kubernetes cluster")
	startCmd.Flags().Bool("dns-proxy", false, "Enable proxy for NAT DNS requests (virtualbox driver only)")
	startCmd.Flags().StringSlice("docker-env", nil, "Environment variables to pass to the Docker daemon. (format: key=value)")
	startCmd.Flags().StringSlice("docker-opt", nil, "Specify arbitrary flags to pass to the Docker daemon. (format: key=value)")
	startCmd.Flags().Bool("download-only", false, "If true, only download and cache files for later use - don't install or start anything.")
	startCmd.Flags().StringP("driver", "d", "", "Driver is one of: virtualbox, kvm2, qemu2, qemu, vmware, none, docker, podman, ssh (defaults to auto-detect)")
	startCmd.Flags().Bool("dry-run", false, "dry-run mode. Validates configuration, but does not mutate system state")
	startCmd.Flags().Bool("embed-certs", false, "if true, will embed the certs in kubeconfig.")
	startCmd.Flags().Bool("enable-default-cni", false, "DEPRECATED: Replaced by --cni=bridge")
	startCmd.Flags().String("extra-config", "", "A set of key=value pairs that describe configuration that may be passed to different components.")
	startCmd.Flags().String("extra-disks", "", "Number of extra disks created and attached to the minikube VM (currently only implemented for hyperkit, kvm2, qemu2, vfkit, and krunkit drivers)")
	startCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for alpha/experimental features.")
	startCmd.Flags().Bool("force", false, "Force minikube to perform possibly dangerous operations")
	startCmd.Flags().Bool("force-systemd", false, "If set, force the container runtime to use systemd as cgroup manager. Defaults to false.")
	startCmd.Flags().StringP("gpus", "g", "", "Allow pods to use your GPUs. Options include: [all,nvidia,amd] (Docker driver with Docker container-runtime only)")
	startCmd.Flags().Bool("ha", false, "Create Highly Available Multi-Control Plane Cluster with a minimum of three control-plane nodes that will also be marked for work.")
	startCmd.Flags().Bool("host-dns-resolver", false, "Enable host resolver for NAT DNS requests (virtualbox driver only)")
	startCmd.Flags().String("host-only-cidr", "", "The CIDR to be used for the minikube VM (virtualbox driver only)")
	startCmd.Flags().String("host-only-nic-type", "", "NIC Type used for host only network. One of Am79C970A, Am79C973, 82540EM, 82543GC, 82545EM, or virtio (virtualbox driver only)")
	startCmd.Flags().String("hyperkit-vpnkit-sock", "", "Location of the VPNKit socket used for networking. If empty, disables Hyperkit VPNKitSock, if 'auto' uses Docker for Mac VPNKit connection, otherwise uses the specified VSock (hyperkit driver only)")
	startCmd.Flags().StringSlice("hyperkit-vsock-ports", nil, "List of guest VSock ports that should be exposed as sockets on the host (hyperkit driver only)")
	startCmd.Flags().String("hyperv-external-adapter", "", "External Adapter on which external switch will be created if no external switch is found. (hyperv driver only)")
	startCmd.Flags().Bool("hyperv-use-external-switch", false, "Whether to use external switch over Default Switch if virtual switch not explicitly specified. (hyperv driver only)")
	startCmd.Flags().String("hyperv-virtual-switch", "", "The hyperv virtual switch name. Defaults to first found. (hyperv driver only)")
	startCmd.Flags().String("image-mirror-country", "", "Country code of the image mirror to be used. Leave empty to use the global one. For Chinese mainland users, set it to cn.")
	startCmd.Flags().String("image-repository", "", "Alternative image repository to pull docker images from. This can be used when you have limited access to gcr.io. Set it to \"auto\" to let minikube decide one for you. For Chinese mainland users, you may use local gcr.io mirrors such as registry.cn-hangzhou.aliyuncs.com/google_containers")
	startCmd.Flags().StringSlice("insecure-registry", nil, "Insecure Docker registries to pass to the Docker daemon.  The default service CIDR range will automatically be added.")
	startCmd.Flags().Bool("install-addons", false, "If set, install addons. Defaults to true.")
	startCmd.Flags().Bool("interactive", false, "Allow user prompts for more information")
	startCmd.Flags().StringSlice("iso-url", nil, "Locations to fetch the minikube ISO from.")
	startCmd.Flags().Bool("keep-context", false, "This will keep the existing kubectl context and will create a minikube context.")
	startCmd.Flags().String("kubernetes-version", "", "The Kubernetes version that the minikube VM will use (ex: v1.2.3, 'stable' for v1.35.0, 'latest' for v1.35.0). Defaults to 'stable'.")
	startCmd.Flags().Bool("kvm-gpu", false, "Enable experimental NVIDIA GPU support in minikube")
	startCmd.Flags().Bool("kvm-hidden", false, "Hide the hypervisor signature from the guest in minikube (kvm2 driver only)")
	startCmd.Flags().String("kvm-network", "", "The KVM default network name. (kvm2 driver only)")
	startCmd.Flags().String("kvm-numa-count", "", "Simulate numa node count in minikube, supported numa node count range is 1-8 (kvm2 driver only)")
	startCmd.Flags().String("kvm-qemu-uri", "", "The KVM QEMU connection URI. (kvm2 driver only)")
	startCmd.Flags().String("listen-address", "", "IP Address to use to expose ports (docker and podman driver only)")
	startCmd.Flags().StringP("memory", "m", "", "Amount of RAM to allocate to Kubernetes (format: <number>[<unit>], where unit = b, k, m or g). Use \"max\" to use the maximum amount of memory. Use \"no-limit\" to not specify a limit (Docker/Podman only)")
	startCmd.Flags().Bool("mount", false, "Kept for backward compatibility, value is ignored.")
	startCmd.Flags().String("mount-9p-version", "", "Specify the 9p version that the mount should use")
	startCmd.Flags().String("mount-gid", "", "Default group id used for the mount")
	startCmd.Flags().String("mount-ip", "", "Specify the ip that the mount should be setup on")
	startCmd.Flags().String("mount-msize", "", "The number of bytes to use for 9p packet payload")
	startCmd.Flags().StringSlice("mount-options", nil, "Additional mount options, such as cache=fscache")
	startCmd.Flags().String("mount-port", "", "Specify the port that the mount should be setup on, where 0 means any free port.")
	startCmd.Flags().String("mount-string", "", "Directory to mount in the guest using format '/host-path:/guest-path'.")
	startCmd.Flags().String("mount-type", "", "Specify the mount filesystem type (supported types: 9p)")
	startCmd.Flags().String("mount-uid", "", "Default user id used for the mount")
	startCmd.Flags().String("namespace", "", "The named space to activate after start")
	startCmd.Flags().String("nat-nic-type", "", "NIC Type used for nat network. One of Am79C970A, Am79C973, 82540EM, 82543GC, 82545EM, or virtio (virtualbox driver only)")
	startCmd.Flags().Bool("native-ssh", false, "Use native Golang SSH client (default true). Set to 'false' to use the command line 'ssh' command when accessing the docker machine. Useful for the machine drivers when they will not start with 'Waiting for SSH'.")
	startCmd.Flags().String("network", "", "network to run minikube with. Used by docker/podman, qemu, kvm, and vfkit drivers. If left empty, minikube will create a new network.")
	startCmd.Flags().String("network-plugin", "", "DEPRECATED: Replaced by --cni")
	startCmd.Flags().StringSlice("nfs-share", nil, "Local folders to share with Guest via NFS mounts (hyperkit driver only)")
	startCmd.Flags().String("nfs-shares-root", "", "Where to root the NFS Shares, defaults to /nfsshares (hyperkit driver only)")
	startCmd.Flags().Bool("no-kubernetes", false, "If set, minikube VM/container will start without starting or configuring Kubernetes. (only works on new clusters)")
	startCmd.Flags().Bool("no-vtx-check", false, "Disable checking for the availability of hardware virtualization before the vm is started (virtualbox driver only)")
	startCmd.Flags().StringP("nodes", "n", "", "The total number of nodes to spin up. Defaults to 1.")
	startCmd.Flags().StringP("output", "o", "", "Format to print stdout in. Options include: [text,json]")
	startCmd.Flags().StringSlice("ports", nil, "List of ports that should be exposed (docker and podman driver only)")
	startCmd.Flags().Bool("preload", false, "If set, download tarball of preloaded images if available to improve start time. Defaults to true.")
	startCmd.Flags().String("preload-source", "", "Which source to download the preload from (valid options: gcs, github, auto). Defaults to auto (try both).")
	startCmd.Flags().String("qemu-firmware-path", "", "Path to the qemu firmware file. Defaults: For Linux, the default firmware location. For macOS, the brew installation location. For Windows, C:\\Program Files\\qemu\\share")
	startCmd.Flags().StringSlice("registry-mirror", nil, "Registry mirrors to pass to the Docker daemon")
	startCmd.Flags().Bool("rosetta", false, "Enable Rosetta to support apps built for Intel processor on a Mac with Apple silicon (vfkit driver only)")
	startCmd.Flags().String("service-cluster-ip-range", "", "The CIDR to be used for service cluster IPs.")
	startCmd.Flags().String("socket-vmnet-client-path", "", "Path to the socket vmnet client binary (QEMU driver only)")
	startCmd.Flags().String("socket-vmnet-path", "", "Path to socket vmnet binary (QEMU driver only)")
	startCmd.Flags().String("ssh-ip-address", "", "IP address (ssh driver only)")
	startCmd.Flags().String("ssh-key", "", "SSH key (ssh driver only)")
	startCmd.Flags().String("ssh-port", "", "SSH port (ssh driver only)")
	startCmd.Flags().String("ssh-user", "", "SSH user (ssh driver only)")
	startCmd.Flags().String("static-ip", "", "Set a static IP for the minikube cluster, the IP must be: private, IPv4, and the last octet must be between 2 and 254, for example 192.168.200.200 (Docker and Podman drivers only)")
	startCmd.Flags().String("subnet", "", "Subnet to be used on kic cluster. If left empty, minikube will choose subnet address, beginning from 192.168.49.0. (docker and podman driver only)")
	startCmd.Flags().String("trace", "", "Send trace events. Options include: [gcp]")
	startCmd.Flags().String("uuid", "", "Provide VM UUID to restore MAC address (hyperkit driver only)")
	startCmd.Flags().Bool("vm", false, "Filter to use only VM Drivers")
	startCmd.Flags().String("vm-driver", "", "DEPRECATED, use `driver` instead.")
	startCmd.Flags().StringSlice("wait", nil, "comma separated list of Kubernetes components to verify and wait for after starting a cluster. defaults to \"apiserver,system_pods\", available options: \"apiserver,system_pods,default_sa,apps_running,node_ready,kubelet,extra\" . other acceptable values are 'all' or 'none', 'true' and 'false'")
	startCmd.Flags().String("wait-timeout", "", "max time to wait per Kubernetes or host to be healthy.")
	startCmd.Flag("enable-default-cni").Hidden = true
	startCmd.Flag("network-plugin").Hidden = true
	startCmd.Flag("vm-driver").Hidden = true
	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).FlagCompletion(carapace.ActionMap{
		"addons":     action.ActionAddons().UniqueList(","),
		"base-image": docker.ActionRepositoryTags(),
		"cni": carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValues("auto", "bridge", "calico", "cilium", "flannel", "kindnet"),
		).ToA(),
		"container-runtime":  carapace.ActionValues("docker", "cri-o", "containerd"),
		"cri-socket":         carapace.ActionFiles(),
		"driver":             carapace.ActionValues("virtualbox", "vmwarefusion", "kvm2", "vmware", "none", "docker", "podman", "ssh", "auto-detect"),
		"host-only-nic-type": carapace.ActionValues("Am79C970A", "Am79C973", "82540EM", "82543GC", "82545EM", "virtio"),
		"hyperkit-vpnkit-sock": carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValues("auto"),
		).ToA(),
		"image-mirror-country": os.ActionLanguages(), // TODO country codes the same?
		"mount-string": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles().NoSpace()
			case 1:
				return carapace.ActionValues("/minikube-host")
			default:
				return carapace.ActionValues()
			}
		}),
		"nat-nic-type":   carapace.ActionValues("Am79C970A", "Am79C973", "82540EM", "82543GC", "82545EM", "virtio"),
		"output":         carapace.ActionValues("text", "json"),
		"preload-source": carapace.ActionValues("gcs", "github", "auto").StyleF(style.ForKeyword),
		"ssh-key":        carapace.ActionFiles(),
		"ssh-user":       os.ActionUsers(),
		"wait": carapace.ActionValues(
			"apiserver",
			"system_pods",
			"default_sa",
			"apps_running",
			"node_ready",
			"kubelet",
			"all",
			"none",
		).UniqueList(","),
	})
}
