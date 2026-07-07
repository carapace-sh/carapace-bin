package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var isCmd = &cobra.Command{
	Use:   "is",
	Short: "Check if paths match type/permission",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isCmd).Standalone()

	isCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	isCmd.Flags().Bool("all", false, "return 0 only if all paths pass")
	isCmd.Flags().BoolS("d", "d", false, "check for directories")
	isCmd.Flags().BoolS("f", "f", false, "check for regular files")
	isCmd.Flags().BoolP("invert", "v", false, "invert the check")
	isCmd.Flags().BoolS("l", "l", false, "check for symbolic links")
	isCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	isCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	isCmd.Flags().StringP("perm", "p", "", "permission check")
	isCmd.Flags().BoolS("q", "q", false, "suppress output")
	isCmd.Flags().Bool("quiet", false, "suppress output")
	isCmd.Flags().BoolS("r", "r", false, "check readable paths")
	isCmd.Flags().StringP("type", "t", "", "type check")
	isCmd.Flags().BoolS("w", "w", false, "check writable paths")
	isCmd.Flags().BoolS("x", "x", false, "check executable paths")
	isCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(isCmd)

	carapace.Gen(isCmd).FlagCompletion(carapace.ActionMap{
		"perm": carapace.ActionValues("read", "write", "exec", "suid", "sgid", "user", "group"),
		"type": carapace.ActionValues("dir", "file", "link", "block", "char", "fifo", "socket"),
	})
}
