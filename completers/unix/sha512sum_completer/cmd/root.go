package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sha512sum",
	Short: "Print or check SHA512 (512-bit) checksums",
	Long:  "https://www.man7.org/linux/man-pages/man1/sha512sum.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("binary", "b", false, "read in binary mode")
	rootCmd.Flags().BoolP("check", "c", false, "read checksums from the FILEs and check them")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("ignore-missing", false, "don't fail or report status for missing files")
	rootCmd.Flags().Bool("quiet", false, "don't print OK for each successfully verified file")
	rootCmd.Flags().Bool("status", false, "don't output anything, status code shows success")
	rootCmd.Flags().Bool("strict", false, "exit non-zero for improperly formatted checksum lines")
	rootCmd.Flags().Bool("tag", false, "create a BSD-style checksum")
	rootCmd.Flags().BoolP("text", "t", false, "read in text mode")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("warn", "w", false, "warn about improperly formatted checksum lines")
	rootCmd.Flags().BoolP("zero", "z", false, "end each output line with NUL")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
