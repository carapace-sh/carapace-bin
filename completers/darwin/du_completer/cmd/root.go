package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "du",
	Short: "display disk usage statistics",
	Long:  "https://keith.github.io/xcode-manpages/du.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Display an entry for each file in a file hierarchy")
	rootCmd.Flags().BoolP("bytes", "b", false, "Display bytes")
	rootCmd.Flags().StringP("depth", "d", "", "Display all totals to the specified depth")
	rootCmd.Flags().BoolP("human-readable", "h", false, "Use unit suffixes: Byte, Kilobyte, Megabyte, Gigabyte, Terabyte and Petabyte")
	rootCmd.Flags().BoolP("inodes", "I", false, "Display the number of inodes in the file system")
	rootCmd.Flags().BoolP("kilobytes", "k", false, "Use 1024-byte (1-Kbyte) blocks")
	rootCmd.Flags().BoolP("links", "l", false, "Calculate the block usage for symbolic links")
	rootCmd.Flags().BoolP("megabytes", "m", false, "Use 1048576-byte (1-Mbyte) blocks")
	rootCmd.Flags().BoolP("no-dereference", "P", false, "Do not follow symbolic links")
	rootCmd.Flags().BoolP("separate-dirs", "S", false, "Display only a grand total for each directory specified")
	rootCmd.Flags().BoolP("si", "H", false, "Use powers of 1000 for -h (SI units)")
	rootCmd.Flags().BoolP("silent", "q", false, "Silence warnings about directories that cannot be read")
	rootCmd.Flags().BoolP("summary", "s", false, "Display only the grand total for each directory specified")
	rootCmd.Flags().StringP("threshold", "t", "", "Display only entries for which the usage is at least the specified threshold")
	rootCmd.Flags().BoolP("total", "c", false, "Display a grand total")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
