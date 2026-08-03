package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "7z",
	Short: "A file archiver with high compression ratio",
	Long:  "https://linux.die.net/man/1/7z",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	addCommonFlags(rootCmd)
	addCommonFlagCompletions(rootCmd)
}

// addCommonFlags registers switches shared across all 7z subcommands.
func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("ai", "ai", "", "include archive filenames")
	cmd.Flags().BoolS("an", "an", false, "disable archive_name field")
	cmd.Flags().StringS("ao", "ao", "", "set overwrite mode")
	cmd.Flags().StringS("ax", "ax", "", "exclude archive filenames")
	cmd.Flags().StringS("bb", "bb", "", "set output log level")
	cmd.Flags().BoolS("bd", "bd", false, "disable progress indicator")
	cmd.Flags().StringS("bs", "bs", "", "set output stream for output/error/progress")
	cmd.Flags().BoolS("bt", "bt", false, "show execution time statistics")
	cmd.Flags().StringS("i", "i", "", "include filenames")
	cmd.Flags().StringS("m", "m", "", "set compression method")
	cmd.Flags().StringS("o", "o", "", "set output directory")
	cmd.Flags().StringS("p", "p", "", "set password")
	cmd.Flags().StringS("r", "r", "", "recurse subdirectories")
	cmd.Flags().StringS("sa", "sa", "", "set archive name mode")
	cmd.Flags().StringS("scc", "scc", "", "set charset for console input/output")
	cmd.Flags().StringS("scrc", "scrc", "", "set hash function")
	cmd.Flags().StringS("scs", "scs", "", "set charset for list files")
	cmd.Flags().BoolS("sdel", "sdel", false, "delete files after compression")
	cmd.Flags().StringS("seml", "seml", "", "send archive by email")
	cmd.Flags().StringS("sfx", "sfx", "", "create SFX archive")
	cmd.Flags().StringS("si", "si", "", "read data from stdin")
	cmd.Flags().StringS("slp", "slp", "", "set large pages mode")
	cmd.Flags().BoolS("slt", "slt", false, "show technical information for l command")
	cmd.Flags().BoolS("snh", "snh", false, "store hard links as links")
	cmd.Flags().BoolS("sni", "sni", false, "store NT security information")
	cmd.Flags().BoolS("snl", "snl", false, "store symbolic links as links")
	cmd.Flags().StringS("sns", "sns", "", "store NTFS alternate streams")
	cmd.Flags().BoolS("so", "so", false, "write data to stdout")
	cmd.Flags().BoolS("spd", "spd", false, "disable wildcard matching for file names")
	cmd.Flags().BoolS("spe", "spe", false, "eliminate duplication of root folder for extract command")
	cmd.Flags().StringS("spf", "spf", "", "use fully qualified file paths")
	cmd.Flags().StringS("ssc", "ssc", "", "set sensitive case mode")
	cmd.Flags().BoolS("sse", "sse", false, "stop archive creating if input file cannot be opened")
	cmd.Flags().BoolS("ssp", "ssp", false, "do not change Last Access Time of source files")
	cmd.Flags().BoolS("ssw", "ssw", false, "compress shared files")
	cmd.Flags().BoolS("stl", "stl", false, "set archive timestamp from the most recently modified file")
	cmd.Flags().StringS("stm", "stm", "", "set CPU thread affinity mask")
	cmd.Flags().StringS("stx", "stx", "", "exclude archive type")
	cmd.Flags().StringS("t", "t", "", "set type of archive")
	cmd.Flags().StringS("u", "u", "", "update options")
	cmd.Flags().StringS("v", "v", "", "create volumes")
	cmd.Flags().StringS("w", "w", "", "set working directory")
	cmd.Flags().StringS("x", "x", "", "exclude filenames")
	cmd.Flags().BoolS("y", "y", false, "assume Yes on all queries")

	cmd.Flag("ai").NoOptDefVal = " "
	cmd.Flag("ai").OptargDelimiter = -1
	cmd.Flag("ao").NoOptDefVal = " "
	cmd.Flag("ao").OptargDelimiter = -1
	cmd.Flag("ax").NoOptDefVal = " "
	cmd.Flag("ax").OptargDelimiter = -1
	cmd.Flag("bb").NoOptDefVal = " "
	cmd.Flag("bb").OptargDelimiter = -1
	cmd.Flag("bs").NoOptDefVal = " "
	cmd.Flag("bs").OptargDelimiter = -1
	cmd.Flag("i").NoOptDefVal = " "
	cmd.Flag("i").OptargDelimiter = -1
	cmd.Flag("m").NoOptDefVal = " "
	cmd.Flag("m").OptargDelimiter = -1
	cmd.Flag("o").NoOptDefVal = " "
	cmd.Flag("o").OptargDelimiter = -1
	cmd.Flag("p").NoOptDefVal = " "
	cmd.Flag("p").OptargDelimiter = -1
	cmd.Flag("r").NoOptDefVal = " "
	cmd.Flag("r").OptargDelimiter = -1
	cmd.Flag("sa").NoOptDefVal = " "
	cmd.Flag("sa").OptargDelimiter = -1
	cmd.Flag("scc").NoOptDefVal = " "
	cmd.Flag("scc").OptargDelimiter = -1
	cmd.Flag("scrc").NoOptDefVal = " "
	cmd.Flag("scrc").OptargDelimiter = -1
	cmd.Flag("scs").NoOptDefVal = " "
	cmd.Flag("scs").OptargDelimiter = -1
	cmd.Flag("seml").NoOptDefVal = " "
	cmd.Flag("seml").OptargDelimiter = -1
	cmd.Flag("sfx").NoOptDefVal = " "
	cmd.Flag("sfx").OptargDelimiter = -1
	cmd.Flag("si").NoOptDefVal = " "
	cmd.Flag("si").OptargDelimiter = -1
	cmd.Flag("slp").NoOptDefVal = " "
	cmd.Flag("slp").OptargDelimiter = -1
	cmd.Flag("sns").NoOptDefVal = " "
	cmd.Flag("sns").OptargDelimiter = -1
	cmd.Flag("spf").NoOptDefVal = " "
	cmd.Flag("spf").OptargDelimiter = -1
	cmd.Flag("ssc").NoOptDefVal = " "
	cmd.Flag("ssc").OptargDelimiter = -1
	cmd.Flag("stm").NoOptDefVal = " "
	cmd.Flag("stm").OptargDelimiter = -1
	cmd.Flag("stx").NoOptDefVal = " "
	cmd.Flag("stx").OptargDelimiter = -1
	cmd.Flag("t").NoOptDefVal = " "
	cmd.Flag("t").OptargDelimiter = -1
	cmd.Flag("u").NoOptDefVal = " "
	cmd.Flag("u").OptargDelimiter = -1
	cmd.Flag("v").NoOptDefVal = " "
	cmd.Flag("v").OptargDelimiter = -1
	cmd.Flag("w").NoOptDefVal = " "
	cmd.Flag("w").OptargDelimiter = -1
	cmd.Flag("x").NoOptDefVal = " "
	cmd.Flag("x").OptargDelimiter = -1
}

// addCommonFlagCompletions registers flag completions for common switches.
func addCommonFlagCompletions(cmd *cobra.Command) {
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"ao":   carapace.ActionValues("a", "s", "t", "u"),
		"bb":   carapace.ActionValues("0", "1", "2", "3"),
		"o":    carapace.ActionDirectories(),
		"r":    carapace.ActionValues("-", "0"),
		"sa":   carapace.ActionValues("a", "e", "s"),
		"scc":  carapace.ActionValues("UTF-8", "WIN", "DOS"),
		"scrc": carapace.ActionValues("CRC32", "CRC64", "SHA1", "SHA256", "XXH64", "*"),
		"scs":  carapace.ActionValues("UTF-8", "UTF-16LE", "UTF-16BE", "WIN", "DOS"),
		"seml": carapace.ActionValues("."),
		"sfx":  carapace.ActionFiles(),
		"slp":  carapace.ActionValues("-"),
		"sns":  carapace.ActionValues("-"),
		"spf":  carapace.ActionValues("2"),
		"ssc":  carapace.ActionValues("-"),
		"t": carapace.ActionValues(
			"7z", "bzip2", "cab", "gzip",
			"iso", "lzma", "lzma86",
			"tar", "udf", "wim", "xz", "zip",
			"zstd", "split",
		),
		"w": carapace.ActionDirectories(),
	})
}
