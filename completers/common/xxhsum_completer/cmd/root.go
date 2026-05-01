package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xxhsum",
	Short: "Create or verify checksums using fast non-cryptographic algorithm xxHash",
	Long:  "https://github.com/Cyan4973/xxHash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("H", "H", "", "select an xxhash algorithm")
	rootCmd.Flags().StringS("b", "b", "", "bench only algorithm variant X")
	rootCmd.Flags().Bool("binary", false, "read in binary mode")
	rootCmd.Flags().BoolP("check", "c", false, "read xxHash checksum from [files] and check them")
	rootCmd.Flags().Bool("filelist", false, "generate hashes for files listed in [files]")
	rootCmd.Flags().Bool("files-from", false, "generate hashes for files listed in [files]")
	rootCmd.Flags().BoolP("help", "h", false, "display a long help page about advanced options")
	rootCmd.Flags().StringS("i", "i", "", "number of times to run the benchmark (default: 3)")
	rootCmd.Flags().Bool("ignore-missing", false, "don't fail or report status for missing files")
	rootCmd.Flags().Bool("little-endian", false, "checksum values use little endian convention (default: big endian)")
	rootCmd.Flags().BoolP("quiet", "q", false, "don't display version header in benchmark mode")
	rootCmd.Flags().Bool("status", false, "don't output anything, status code shows success")
	rootCmd.Flags().Bool("strict", false, "exit non-zero for improperly formatted lines in [files]")
	rootCmd.Flags().Bool("tag", false, "produce BSD-style checksum lines")
	rootCmd.Flags().BoolP("version", "V", false, "display version information")
	rootCmd.Flags().Bool("warn", false, "warn about improperly formatted lines in [files]")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"H": carapace.ActionValuesDescribed(
			"0", "XXH32",
			"1", "XXH64 (default)",
			"2", "XXH128 (also called XXH3_128bits)",
			"3", "XXH3 (also called XXH3_64bits)",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
