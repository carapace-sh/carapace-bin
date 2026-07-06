package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/qemu-img_completer/cmd/action"
	"github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Run a simple image benchmark",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(benchCmd).Standalone()

	benchCmd.Flags().StringP("aio", "i", "", "async-io backend (threads, native, io_uring)")
	benchCmd.Flags().StringP("buffer-size", "s", "", "size of each I/O request")
	benchCmd.Flags().StringP("cache", "t", "", "cache mode for FILE (default: writeback)")
	benchCmd.Flags().StringP("count", "c", "", "number of I/O requests to perform")
	benchCmd.Flags().StringP("depth", "d", "", "number of requests to perform in parallel")
	benchCmd.Flags().String("flush-interval", "", "issue flush after this number of requests")
	benchCmd.Flags().BoolP("force-share", "U", false, "open image in shared mode for concurrent access")
	benchCmd.Flags().StringP("format", "f", "", "specify FILE format explicitly")
	benchCmd.Flags().Bool("image-opts", false, "indicates that FILE is a complete image specification")
	benchCmd.Flags().BoolP("native", "n", false, "use native AIO backend if possible")
	benchCmd.Flags().Bool("no-drain", false, "do not wait when flushing pending requests")
	benchCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	benchCmd.Flags().StringP("offset", "o", "", "start first request at this OFFSET")
	benchCmd.Flags().String("pattern", "", "write this pattern byte instead of zero")
	benchCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	benchCmd.Flags().StringP("step-size", "S", "", "each next request offset increment")
	benchCmd.Flags().BoolP("write", "w", false, "perform write test (default is read)")
	rootCmd.AddCommand(benchCmd)

	carapace.Gen(benchCmd).FlagCompletion(carapace.ActionMap{
		"aio":    action.ActionAioModes(),
		"cache":  action.ActionCacheModes(),
		"format": action.ActionImageFormats(),
	})

	carapace.Gen(benchCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
