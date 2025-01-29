package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "podman [options]",
	Short: "Simple management tool for pods, containers and images",
	Long:  "https://podman.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("cgroup-manager", "", "Cgroup manager to use (\"cgroupfs\"|\"systemd\")")
	rootCmd.Flags().String("config", "", "Location of authentication config file")
	rootCmd.PersistentFlags().String("conmon", "", "Path of the conmon binary")
	rootCmd.Flags().StringP("connection", "c", "", "Connection to use for remote Podman service (CONTAINER_CONNECTION)")
	rootCmd.Flags().String("context", "", "Name of the context to use to connect to the daemon (This flag is a NOOP and provided solely for scripting compatibility.)")
	rootCmd.PersistentFlags().String("cpu-profile", "", "Path for the cpu-profiling results")
	rootCmd.PersistentFlags().String("db-backend", "", "Database backend to use")
	rootCmd.Flags().BoolP("debug", "D", false, "Docker compatibility, force setting of log-level")
	rootCmd.PersistentFlags().String("default-mounts-file", "", "Path to default mounts file")
	rootCmd.PersistentFlags().String("events-backend", "", "Events backend to use (\"file\"|\"journald\"|\"none\")")
	rootCmd.PersistentFlags().Bool("help", false, "Help for podman")
	rootCmd.PersistentFlags().StringSlice("hooks-dir", []string{}, "Set the OCI hooks directory path (may be set multiple times)")
	rootCmd.Flags().StringP("host", "H", "", "Used for Docker compatibility")
	rootCmd.Flags().String("identity", "", "path to SSH identity file, (CONTAINER_SSHKEY)")
	rootCmd.PersistentFlags().String("imagestore", "", "Path to the 'image store', different from 'graph root', use this to split storing the image into a separate 'image store', see 'man containers-storage.conf' for details")
	rootCmd.PersistentFlags().String("log-level", "", "Log messages above specified level (trace, debug, info, warn, warning, error, fatal, panic)")
	rootCmd.PersistentFlags().String("max-workers", "", "The maximum number of workers for parallel operations")
	rootCmd.PersistentFlags().String("memory-profile", "", "Path for the memory-profiling results")
	rootCmd.Flags().StringSlice("module", []string{}, "Load the containers.conf(5) module")
	rootCmd.PersistentFlags().String("namespace", "", "Set the libpod namespace, used to create separate views of the containers and pods on the system")
	rootCmd.PersistentFlags().String("network-backend", "", "Network backend to use (\"cni\"|\"netavark\")")
	rootCmd.PersistentFlags().String("network-cmd-path", "", "Path to the command for configuring the network")
	rootCmd.PersistentFlags().String("network-config-dir", "", "Path of the configuration directory for networks")
	rootCmd.Flags().Bool("noout", false, "do not output to stdout")
	rootCmd.Flags().String("out", "", "Send output (stdout) from podman to a file")
	rootCmd.PersistentFlags().StringSlice("pull-option", []string{}, "Specify an option to change how the image is pulled")
	rootCmd.PersistentFlags().String("registries-conf", "", "Path to a registries.conf to use for image processing")
	rootCmd.Flags().BoolP("remote", "r", false, "Access remote Podman service")
	rootCmd.PersistentFlags().String("root", "", "Path to the graph root directory where images, containers, etc. are stored")
	rootCmd.PersistentFlags().String("runroot", "", "Path to the 'run directory' where all state information is stored")
	rootCmd.PersistentFlags().String("runtime", "", "Path to the OCI-compatible binary used to run containers.")
	rootCmd.PersistentFlags().StringSlice("runtime-flag", []string{}, "add global flags for the container runtime")
	rootCmd.Flags().String("ssh", "", "define the ssh mode")
	rootCmd.PersistentFlags().String("storage-driver", "", "Select which storage driver is used to manage storage of images and containers")
	rootCmd.PersistentFlags().StringSlice("storage-opt", []string{}, "Used to pass an option to the storage driver")
	rootCmd.PersistentFlags().Bool("syslog", false, "Output logging information to syslog as well as the console (default false)")
	rootCmd.PersistentFlags().String("tmpdir", "", "Path to the tmp directory for libpod state content.")
	rootCmd.PersistentFlags().Bool("trace", false, "Enable opentracing output (default false)")
	rootCmd.PersistentFlags().Bool("transient-store", false, "Enable transient container storage")
	rootCmd.Flags().String("url", "", "URL to access Podman service (CONTAINER_HOST)")
	rootCmd.PersistentFlags().String("volumepath", "", "Path to the volume directory in which volume data is stored")
	rootCmd.Flag("context").Hidden = true
	rootCmd.Flag("cpu-profile").Hidden = true
	rootCmd.Flag("db-backend").Hidden = true
	rootCmd.Flag("debug").Hidden = true
	rootCmd.Flag("default-mounts-file").Hidden = true
	rootCmd.Flag("host").Hidden = true
	rootCmd.Flag("max-workers").Hidden = true
	rootCmd.Flag("memory-profile").Hidden = true
	rootCmd.Flag("namespace").Hidden = true
	rootCmd.Flag("network-backend").Hidden = true
	rootCmd.Flag("noout").Hidden = true
	rootCmd.Flag("pull-option").Hidden = true
	rootCmd.Flag("registries-conf").Hidden = true
	rootCmd.Flag("trace").Hidden = true
}
