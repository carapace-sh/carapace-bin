package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-nbd",
	Short: "QEMU Disk Network Block Device Utility",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("aio", "", "set AIO mode (native, io_uring or threads)")
	rootCmd.Flags().BoolP("allocation-depth", "A", false, "expose the allocation depth")
	rootCmd.Flags().StringP("bind", "b", "", "interface to bind to (default `0.0.0.0')")
	rootCmd.Flags().StringP("bitmap", "B", "", "expose a persistent dirty bitmap")
	rootCmd.Flags().String("cache", "", "set cache mode used to access the disk image")
	rootCmd.Flags().StringP("connect", "c", "", "connect FILE to the local NBD device DEV")
	rootCmd.Flags().StringP("description", "D", "", "export a human-readable description")
	rootCmd.Flags().String("detect-zeroes", "", "set detect-zeroes mode (off, on, unmap)")
	rootCmd.Flags().String("discard", "", "set discard mode (ignore, unmap)")
	rootCmd.Flags().BoolP("disconnect", "d", false, "disconnect the specified device")
	rootCmd.Flags().StringP("export-name", "x", "", "expose export by name (default is empty string)")
	rootCmd.Flags().Bool("fork", false, "fork off the server process and exit the parent once the server is running")
	rootCmd.Flags().StringP("format", "f", "", "set image format (raw, qcow2, ...)")
	rootCmd.Flags().String("handshake-limit", "", "limit client's handshake to N seconds (default 10)")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().Bool("image-opts", false, "treat FILE as a full set of image options")
	rootCmd.Flags().BoolP("list", "L", false, "list exports available from another NBD server")
	rootCmd.Flags().StringP("load-snapshot", "l", "", "load an internal snapshot inside FILE")
	rootCmd.Flags().BoolP("nocache", "n", false, "disable host cache")
	rootCmd.Flags().String("object", "", "define an object such as 'secret' for providing passwords and/or encryption keys")
	rootCmd.Flags().StringP("offset", "o", "", "offset into the image")
	rootCmd.Flags().BoolP("persistent", "t", false, "don't exit on the last connection")
	rootCmd.Flags().String("pid-file", "", "store the server's process ID in the given file")
	rootCmd.Flags().StringP("port", "p", "", "port to listen on (default `10809')")
	rootCmd.Flags().BoolP("read-only", "r", false, "export read-only")
	rootCmd.Flags().StringP("shared", "e", "", "device can be shared by NUM clients (default '1')")
	rootCmd.Flags().BoolP("snapshot", "s", false, "use FILE as an external snapshot")
	rootCmd.Flags().StringP("socket", "k", "", "path to the unix socket")
	rootCmd.Flags().String("tls-authz", "", "use id of an earlier --object to provide authorization")
	rootCmd.Flags().String("tls-creds", "", "use id of an earlier --object to provide TLS")
	rootCmd.Flags().String("tls-hostname", "", "override hostname used to check x509 certificate")
	rootCmd.Flags().StringP("trace", "T", "", "specify tracing options")
	rootCmd.Flags().BoolP("verbose", "v", false, "display extra debugging information")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"aio":           qemu.ActionAioModes(),
		"cache":         qemu.ActionCacheModes(),
		"connect":       carapace.ActionFiles("/dev/nbd*"),
		"detect-zeroes": carapace.ActionValues("off", "on", "unmap"),
		"discard":       carapace.ActionValues("ignore", "unmap"),
		"format":        qemu.ActionImageFormats(),
		"pid-file":      carapace.ActionFiles(),
		"socket":        carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
