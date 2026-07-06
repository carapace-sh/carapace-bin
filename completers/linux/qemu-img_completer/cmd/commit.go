package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/qemu-img_completer/cmd/action"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit image to its backing file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().StringP("base", "b", "", "image in the backing chain to commit change to")
	commitCmd.Flags().StringP("cache", "t", "", "image cache mode (default: writeback)")
	commitCmd.Flags().BoolP("drop", "d", false, "skip emptying FILE on completion")
	commitCmd.Flags().StringP("format", "f", "", "specify FILE image format explicitly")
	commitCmd.Flags().Bool("image-opts", false, "treat FILE as an option string")
	commitCmd.Flags().String("object", "", "defines QEMU user-creatable object")
	commitCmd.Flags().BoolP("progress", "p", false, "display progress information")
	commitCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	commitCmd.Flags().StringP("rate-limit", "r", "", "I/O rate limit, in bytes per second")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).FlagCompletion(carapace.ActionMap{
		"cache":  action.ActionCacheModes(),
		"format": action.ActionImageFormats(),
	})

	carapace.Gen(commitCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
