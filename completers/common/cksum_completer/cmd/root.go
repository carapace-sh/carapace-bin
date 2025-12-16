package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cksum",
	Short: "compute and verify file checksums",
	Long:  "https://www.man7.org/linux/man-pages/man1/cksum.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("algorithm", "a", "", "select the digest type to use")
	rootCmd.Flags().Bool("base64", false, "emit base64-encoded digests, not hexadecimal")
	rootCmd.Flags().BoolP("check", "c", false, "read checksums from the FILEs and check them")
	rootCmd.Flags().Bool("debug", false, "indicate which implementation used")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("ignore-missing", false, "don't fail or report status for missing files")
	rootCmd.Flags().StringP("length", "l", "", "digest length in bits; must not exceed the max for the blake2 algorithm and must be a multiple of 8")
	rootCmd.Flags().Bool("quiet", false, "don't print OK for each successfully verified file")
	rootCmd.Flags().Bool("raw", false, "emit a raw binary digest, not hexadecimal")
	rootCmd.Flags().Bool("status", false, "don't output anything, status code shows success")
	rootCmd.Flags().Bool("strict", false, "exit non-zero for improperly formatted checksum lines")
	rootCmd.Flags().Bool("tag", false, "create a BSD-style checksum (the default)")
	rootCmd.Flags().Bool("untagged", false, "create a reversed style checksum, without digest type")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("warn", "w", false, "warn about improperly formatted checksum lines")
	rootCmd.Flags().BoolP("zero", "z", false, "end each output line with NUL, not newline, and disable file name escaping")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"algorithm": carapace.ActionValuesDescribed(
			"sysv", "equivalent to sum -s",
			"bsd", "equivalent to sum -r",
			"crc", "equivalent to cksum",
			"crc32b", "only available through cksum",
			"md5", "equivalent to md5sum",
			"sha1", "equivalent to sha1sum",
			"sha224", "equivalent to sha224sum",
			"sha256", "equivalent to sha256sum",
			"sha384", "equivalent to sha384sum",
			"sha512", "equivalent to sha512sum",
			"blake2b", "equivalent to b2sum",
			"sm3", "only available through cksum",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
