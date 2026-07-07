package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-io",
	Short: "QEMU disk exerciser",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("aio", "i", "", "use AIO mode (threads, native or io_uring)")
	rootCmd.Flags().StringP("cache", "t", "", "use the given cache mode for the image")
	rootCmd.Flags().StringP("cmd", "c", "", "execute command with its arguments from the given string")
	rootCmd.Flags().BoolP("copy-on-read", "C", false, "enable copy-on-read")
	rootCmd.Flags().StringP("discard", "d", "", "use the given discard mode for the image")
	rootCmd.Flags().BoolP("force-share", "U", false, "force shared permissions")
	rootCmd.Flags().StringP("format", "f", "", "specifies the block driver to use")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().Bool("image-opts", false, "treat file as option string")
	rootCmd.Flags().BoolP("misalign", "m", false, "misalign allocations for O_DIRECT")
	rootCmd.Flags().BoolP("native-aio", "k", false, "use kernel AIO implementation (Linux only, prefer use of -i)")
	rootCmd.Flags().BoolP("nocache", "n", false, "disable host cache, short for -t none")
	rootCmd.Flags().String("object", "", "define an object such as 'secret' for passwords and/or encryption keys")
	rootCmd.Flags().BoolP("read-only", "r", false, "export read-only")
	rootCmd.Flags().BoolP("snapshot", "s", false, "use snapshot file")
	rootCmd.Flags().StringP("trace", "T", "", "specify tracing options")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"aio":     qemu.ActionAioModes(),
		"cache":   qemu.ActionCacheModes(),
		"discard": carapace.ActionValues("ignore", "unmap"),
		"format":  qemu.ActionImageFormats(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
