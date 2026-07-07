package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mtimeCmd = &cobra.Command{
	Use:   "mtime",
	Short: "Return last modification time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mtimeCmd).Standalone()

	mtimeCmd.Flags().BoolS("R", "R", false, "print relative to now")
	mtimeCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	mtimeCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	mtimeCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	mtimeCmd.Flags().BoolS("q", "q", false, "suppress output")
	mtimeCmd.Flags().Bool("quiet", false, "suppress output")
	mtimeCmd.Flags().Bool("relative", false, "print relative to now")
	mtimeCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(mtimeCmd)
}
