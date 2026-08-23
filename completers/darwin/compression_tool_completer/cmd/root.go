package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compression_tool",
	Short: "encode/decode files using the Compression library",
	Long:  "https://keith.github.io/xcode-manpages/compression_tool.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("A", "", "Enable block compression with the specified algorithm")
	rootCmd.Flags().String("a", "", "Set the compression algorithm")
	rootCmd.Flags().String("b", "", "Set block size for block compression")
	rootCmd.Flags().Bool("decode", false, "Decode (uncompress) the input")
	rootCmd.Flags().Bool("encode", false, "Encode (compress) the input")
	rootCmd.Flags().Bool("h", false, "Print usage and exit")
	rootCmd.Flags().String("i", "", "Input file")
	rootCmd.Flags().String("o", "", "Output file")
	rootCmd.Flags().String("t", "", "Set the number of worker threads")
	rootCmd.Flags().Bool("v", false, "Increase verbosity")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"A": carapace.ActionValues("zlib", "lzma", "lzfse", "lz4"),
		"a": carapace.ActionValues("zlib", "lzma", "lzfse", "lz4", "lz4_raw"),
		"i": carapace.ActionFiles(),
		"o": carapace.ActionFiles(),
	})
}