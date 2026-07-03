package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xz",
	Short: "Compress or decompress .xz and .lzma files",
	Long:  "https://linux.die.net/man/1/xz",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "compression preset")
	rootCmd.Flags().BoolS("1", "1", false, "compression preset")
	rootCmd.Flags().BoolS("2", "2", false, "compression preset")
	rootCmd.Flags().BoolS("3", "3", false, "compression preset")
	rootCmd.Flags().BoolS("4", "4", false, "compression preset")
	rootCmd.Flags().BoolS("5", "5", false, "compression preset")
	rootCmd.Flags().BoolS("6", "6", false, "compression preset")
	rootCmd.Flags().BoolS("7", "7", false, "compression preset")
	rootCmd.Flags().BoolS("8", "8", false, "compression preset")
	rootCmd.Flags().BoolS("9", "9", false, "compression preset")
	rootCmd.Flags().String("arm", "", "ARM BCJ filter")
	rootCmd.Flags().String("arm64", "", "ARM64 BCJ filter")
	rootCmd.Flags().String("armthumb", "", "ARM-Thumb BCJ filter")
	rootCmd.Flags().String("block-list", "", "start a new .xz block after given comma-separated intervals of uncompressed data")
	rootCmd.Flags().String("block-size", "", "start a new .xz block after every SIZE bytes of input")
	rootCmd.Flags().StringP("check", "C", "", "integrity check type")
	rootCmd.Flags().BoolP("compress", "z", false, "force compression")
	rootCmd.Flags().BoolP("decompress", "d", false, "force decompression")
	rootCmd.Flags().String("delta", "", "delta filter")
	rootCmd.Flags().BoolP("extreme", "e", false, "try to improve compression ratio by using more CPU time;")
	rootCmd.Flags().String("files", "", "read filenames to process from FILE; if FILE is omitted, filenames are read from standard input")
	rootCmd.Flags().String("files0", "", "like --files but use the null character as terminator")
	rootCmd.Flags().String("filters", "", "set the filter chain using the liblzma filter string syntax")
	rootCmd.Flags().Bool("filters-help", false, "display information about the liblzma filter string syntax and exit")
	rootCmd.Flags().String("filters1", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("filters2", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("filters3", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("filters4", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("filters5", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("filters6", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("filters7", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("filters8", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("filters9", "", "set additional filter chain to use with --block-list")
	rootCmd.Flags().String("flush-timeout", "", "flush pending data if more than NUM milliseconds has passed since previous flush")
	rootCmd.Flags().BoolP("force", "f", false, "force overwrite of output file and (de)compress links")
	rootCmd.Flags().StringP("format", "F", "", "file format to encode or decode")
	rootCmd.Flags().BoolP("help", "h", false, "display this short help and exit")
	rootCmd.Flags().String("ia64", "", "IA-64 (Itanium) BCJ filter")
	rootCmd.Flags().Bool("ignore-check", false, "don't verify the integrity check when decompressing")
	rootCmd.Flags().Bool("info-memory", false, "display the total amount of RAM and the currently active memory usage limits, and exit")
	rootCmd.Flags().BoolP("keep", "k", false, "keep (don't delete) input files")
	rootCmd.Flags().BoolP("list", "l", false, "list information about .xz files")
	rootCmd.Flags().BoolP("long-help", "H", false, "display the long help (lists also the advanced options)")
	rootCmd.Flags().String("lzma1", "", "LZMA1 filter")
	rootCmd.Flags().String("lzma2", "", "LZMA2 filter")
	rootCmd.Flags().StringP("memlimit", "M", "", "set memory usage limit for compression, decompression, threaded decompression, or all of these")
	rootCmd.Flags().String("memlimit-compress", "", "set memory usage limit for compression")
	rootCmd.Flags().String("memlimit-decompress", "", "set memory usage limit for decompression")
	rootCmd.Flags().String("memlimit-mt-decompress", "", "set memory usage limit for threaded decompression")
	rootCmd.Flags().Bool("no-adjust", false, "if compression settings exceed the memory usage limit, give an error instead of adjusting the settings downwards")
	rootCmd.Flags().Bool("no-sparse", false, "do not create sparse files when decompressing")
	rootCmd.Flags().Bool("no-sync", false, "don't synchronize the output file to the storage device before removing the input file")
	rootCmd.Flags().BoolP("no-warn", "Q", false, "make warnings not affect the exit status")
	rootCmd.Flags().String("powerpc", "", "PowerPC BCJ filter (big endian only)")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress warnings; specify twice to suppress errors too")
	rootCmd.Flags().String("riscv", "", "RISC-V BCJ filter")
	rootCmd.Flags().Bool("robot", false, "use machine-parsable messages (useful for scripts)")
	rootCmd.Flags().Bool("single-stream", false, "decompress only the first stream, and silently ignore possible remaining input data")
	rootCmd.Flags().String("sparc", "", "SPARC BCJ filter")
	rootCmd.Flags().BoolP("stdout", "c", false, "write to standard output and don't delete input files")
	rootCmd.Flags().StringP("suffix", "S", "", "use the suffix '.SUF' on compressed files")
	rootCmd.Flags().BoolP("test", "t", false, "test compressed file integrity")
	rootCmd.Flags().StringP("threads", "T", "", "use at most NUM threads; the default is 0 which uses as many threads as there are processor cores")
	rootCmd.Flags().BoolP("verbose", "v", false, "be verbose; specify twice for even more verbose")
	rootCmd.Flags().BoolP("version", "V", false, "display the version number and exit")
	rootCmd.Flags().String("x86", "", "x86 BCJ filter (32-bit and 64-bit)")

	rootCmd.Flag("files").NoOptDefVal = ""
	rootCmd.Flag("files0").NoOptDefVal = ""
	rootCmd.Flag("arm").NoOptDefVal = ""
	rootCmd.Flag("arm64").NoOptDefVal = ""
	rootCmd.Flag("armthumb").NoOptDefVal = ""
	rootCmd.Flag("delta").NoOptDefVal = ""
	rootCmd.Flag("ia64").NoOptDefVal = ""
	rootCmd.Flag("lzma1").NoOptDefVal = ""
	rootCmd.Flag("lzma2").NoOptDefVal = ""
	rootCmd.Flag("powerpc").NoOptDefVal = ""
	rootCmd.Flag("riscv").NoOptDefVal = ""
	rootCmd.Flag("sparc").NoOptDefVal = ""
	rootCmd.Flag("x86").NoOptDefVal = ""

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"check":  carapace.ActionValues("none", "crc32", "crc64", "sha256"),
		"files":  carapace.ActionFiles(),
		"files0": carapace.ActionFiles(),
		"format": carapace.ActionValues("auto", "xz", "lzma", "lzip", "raw"),
		"suffix": carapace.ActionValues(".xz", ".lzma", ".lzip"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
