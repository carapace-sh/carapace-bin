package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Return paths that match type/permission checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filterCmd).Standalone()

	filterCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	filterCmd.Flags().Bool("all", false, "return 0 only if all paths pass")
	filterCmd.Flags().BoolS("d", "d", false, "filter for directories")
	filterCmd.Flags().BoolS("f", "f", false, "filter for regular files")
	filterCmd.Flags().BoolP("invert", "v", false, "invert the filter")
	filterCmd.Flags().BoolS("l", "l", false, "filter for symbolic links")
	filterCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	filterCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	filterCmd.Flags().StringP("perm", "p", "", "permission filter")
	filterCmd.Flags().BoolS("q", "q", false, "suppress output")
	filterCmd.Flags().Bool("quiet", false, "suppress output")
	filterCmd.Flags().BoolS("r", "r", false, "filter for readable paths")
	filterCmd.Flags().StringP("type", "t", "", "type filter")
	filterCmd.Flags().BoolS("w", "w", false, "filter for writable paths")
	filterCmd.Flags().BoolS("x", "x", false, "filter for executable paths")
	filterCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(filterCmd)

	carapace.Gen(filterCmd).FlagCompletion(carapace.ActionMap{
		"perm": carapace.ActionValues("read", "write", "exec", "suid", "sgid", "user", "group"),
		"type": carapace.ActionValues("dir", "file", "link", "block", "char", "fifo", "socket"),
	})
}
