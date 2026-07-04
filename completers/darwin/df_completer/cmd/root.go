package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "df",
	Short: "display free disk space",
	Long:  "https://keith.github.io/xcode-manpages/df.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("512-blocks", "b", false, "Use 512-byte blocks")
	rootCmd.Flags().BoolS("I", "I", false, "Include statistics on the number of free inodes")
	rootCmd.Flags().BoolP("all", "a", false, "Show all mount points, including those that were mounted with MNT_IGNORE flag")
	rootCmd.Flags().BoolP("gigabytes", "g", false, "Use 1073741824-byte (1-Gbyte) blocks")
	rootCmd.Flags().BoolP("human-readable", "h", false, "Use unit suffixes: Byte, Kilobyte, Megabyte, Gigabyte, Terabyte and Petabyte")
	rootCmd.Flags().BoolP("inode", "i", false, "Show the number of inodes, number of used inodes, and number of free inodes")
	rootCmd.Flags().BoolP("kilobytes", "k", false, "Use 1024-byte (1-Kbyte) blocks, rather than the default")
	rootCmd.Flags().BoolP("local", "l", false, "Only show local filesystems")
	rootCmd.Flags().BoolP("megabytes", "m", false, "Use 1048576-byte (1-Mbyte) blocks, rather than the default")
	rootCmd.Flags().BoolP("no-update", "n", false, "Do not update the cached statistics for any mounted filesystems")
	rootCmd.Flags().BoolP("portability", "P", false, "Use 512-byte blocks and a strict POSIX format")
	rootCmd.Flags().BoolP("si", "H", false, "Use powers of 1000 for -h (SI units)")
	rootCmd.Flags().BoolP("type", "t", false, "Only show filesystems of the specified type")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
