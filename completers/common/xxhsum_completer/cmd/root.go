package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xxhsum",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H0", "H0", false, "select XXH32")
	rootCmd.Flags().BoolS("H1", "H1", false, "select XXH64 (default)")
	rootCmd.Flags().BoolS("H2", "H2", false, "select XXH128 (also called XXH3_128bits)")
	rootCmd.Flags().BoolS("H3", "H3", false, "select XXH3 (also called XXH3_64bits)")
	rootCmd.Flags().BoolS("b", "b", false, "run benchmark")
	rootCmd.Flags().BoolS("bX", "bX", false, "bench only algorithm variant X")
	rootCmd.Flags().Bool("binary", false, "read in binary mode")
	rootCmd.Flags().BoolP("check", "c", false, "read xxHash checksum from [files] and check them")
	rootCmd.Flags().Bool("filelist", false, "generate hashes for files listed in [files]")
	rootCmd.Flags().Bool("files-from", false, "generate hashes for files listed in [files]")
	rootCmd.Flags().BoolP("help", "h", false, "display a long help page about advanced options")
	rootCmd.Flags().BoolS("iX", "iX", false, "number of times to run the benchmark (default: 3)")
	rootCmd.Flags().Bool("ignore-missing", false, "don't fail or report status for missing files")
	rootCmd.Flags().Bool("little-endian", false, "checksum values use little endian convention (default: big endian)")
	rootCmd.Flags().BoolP("quiet", "q", false, "don't display version header in benchmark mode")
	rootCmd.Flags().Bool("status", false, "don't output anything, status code shows success")
	rootCmd.Flags().Bool("strict", false, "exit non-zero for improperly formatted lines in [files]")
	rootCmd.Flags().Bool("tag", false, "produce BSD-style checksum lines")
	rootCmd.Flags().BoolP("version", "V", false, "display version information")
	rootCmd.Flags().Bool("warn", false, "warn about improperly formatted lines in [files]")
}
