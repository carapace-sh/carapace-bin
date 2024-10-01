package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pigz",
	Short: "compress or expand files",
	Long:  "https://linux.die.net/man/1/pigz",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "Compression level 0")
	rootCmd.Flags().BoolS("1", "1", false, "Compression level 1")
	rootCmd.Flags().BoolS("11", "11", false, "Compression level 11, zopfli")
	rootCmd.Flags().BoolS("2", "2", false, "Compression level 2")
	rootCmd.Flags().BoolS("3", "3", false, "Compression level 3")
	rootCmd.Flags().BoolS("4", "4", false, "Compression level 4")
	rootCmd.Flags().BoolS("5", "5", false, "Compression level 5")
	rootCmd.Flags().BoolS("6", "6", false, "Compression level 6")
	rootCmd.Flags().BoolS("7", "7", false, "Compression level 7")
	rootCmd.Flags().BoolS("8", "8", false, "Compression level 8")
	rootCmd.Flags().BoolS("9", "9", false, "Compression level 9")
	rootCmd.Flags().StringP("alias", "A", "", "Use xxx as the name for any --zip entry from stdin")
	rootCmd.Flags().Bool("best", false, "Compression levels 9")
	rootCmd.Flags().StringP("blocksize", "b", "", "Set compression block size to mmmK (default 128K)")
	rootCmd.Flags().StringP("comment", "C", "", "Put comment ccc in the gzip or zip header")
	rootCmd.Flags().BoolP("decompress", "d", false, "Decompress the compressed input")
	rootCmd.Flags().Bool("fast", false, "Compression levels 1")
	rootCmd.Flags().BoolP("first", "F", false, "Do iterations first, before block split for -11")
	rootCmd.Flags().BoolP("force", "f", false, "Force overwrite, compress .gz, links, and to terminal")
	rootCmd.Flags().BoolP("help", "h", false, "Display a help screen and quit")
	rootCmd.Flags().BoolP("huffman", "H", false, "Use only Huffman coding for compression")
	rootCmd.Flags().BoolP("independent", "i", false, "Compress blocks independently for damage recovery")
	rootCmd.Flags().StringP("iterations", "I", "", "Number of iterations for -11 optimization")
	rootCmd.Flags().BoolP("keep", "k", false, "Do not delete original file after processing")
	rootCmd.Flags().BoolP("license", "L", false, "Display the pigz license and quit")
	rootCmd.Flags().BoolP("list", "l", false, "List the contents of the compressed input")
	rootCmd.Flags().StringP("maxsplits", "J", "", "Maximum number of split blocks for -11")
	rootCmd.Flags().BoolP("name", "N", false, "Store or restore file name and mod time")
	rootCmd.Flags().BoolP("no-name", "n", false, "Do not store or restore file name or mod time")
	rootCmd.Flags().BoolP("no-time", "m", false, "Do not store or restore mod time")
	rootCmd.Flags().BoolP("oneblock", "O", false, "Do not split into smaller blocks for -11")
	rootCmd.Flags().StringP("processes", "p", "", "Allow up to n compression threads (default is the")
	rootCmd.Flags().BoolP("quiet", "q", false, "Print no messages, even on error")
	rootCmd.Flags().BoolP("recursive", "r", false, "Process the contents of all subdirectories")
	rootCmd.Flags().BoolP("rle", "U", false, "Use run-length encoding for compression")
	rootCmd.Flags().BoolP("rsyncable", "R", false, "Input-determined block locations for rsync")
	rootCmd.Flags().BoolP("stdout", "c", false, "Write all processed output to stdout (won't delete)")
	rootCmd.Flags().StringP("suffix", "S", "", "Use suffix .sss instead of .gz (for compression)")
	rootCmd.Flags().BoolP("synchronous", "Y", false, "Force output file write to permanent storage")
	rootCmd.Flags().BoolP("test", "t", false, "Test the integrity of the compressed input")
	rootCmd.Flags().BoolP("time", "M", false, "Store or restore mod time")
	rootCmd.Flags().BoolP("verbose", "v", false, "Provide more verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Show the version of pigz")
	rootCmd.Flags().BoolP("zip", "K", false, "Compress to PKWare zip (.zip) single entry format")
	rootCmd.Flags().BoolP("zlib", "z", false, "Compress to zlib (.zz) instead of gzip format")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionPositional(rootCmd),
	)
}
