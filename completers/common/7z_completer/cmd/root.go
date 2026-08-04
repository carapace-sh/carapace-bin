package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

	rootCmd.PersistentFlags().StringS("ai", "ai", "", "include archive filenames")
	rootCmd.PersistentFlags().BoolS("an", "an", false, "disable archive_name field")
	rootCmd.PersistentFlags().StringS("ao", "ao", "", "set overwrite mode")
	rootCmd.PersistentFlags().StringS("ax", "ax", "", "exclude archive filenames")
	rootCmd.PersistentFlags().StringS("bb", "bb", "", "set output log level")
	rootCmd.PersistentFlags().BoolS("bd", "bd", false, "disable progress indicator")
	rootCmd.PersistentFlags().StringS("bs", "bs", "", "set output stream for output/error/progress")
	rootCmd.PersistentFlags().BoolS("bt", "bt", false, "show execution time statistics")
	rootCmd.PersistentFlags().StringS("i", "i", "", "include filenames")
	rootCmd.PersistentFlags().StringS("m", "m", "", "set compression method")
	rootCmd.PersistentFlags().StringS("o", "o", "", "set output directory")
	rootCmd.PersistentFlags().StringS("p", "p", "", "set password")
	rootCmd.PersistentFlags().StringS("r", "r", "", "recurse subdirectories")
	rootCmd.PersistentFlags().StringS("sa", "sa", "", "set archive name mode")
	rootCmd.PersistentFlags().StringS("scc", "scc", "", "set charset for console input/output")
	rootCmd.PersistentFlags().StringS("scrc", "scrc", "", "set hash function")
	rootCmd.PersistentFlags().StringS("scs", "scs", "", "set charset for list files")
	rootCmd.PersistentFlags().BoolS("sdel", "sdel", false, "delete files after compression")
	rootCmd.PersistentFlags().StringS("seml", "seml", "", "send archive by email")
	rootCmd.PersistentFlags().StringS("sfx", "sfx", "", "create SFX archive")
	rootCmd.PersistentFlags().StringS("si", "si", "", "read data from stdin")
	rootCmd.PersistentFlags().StringS("slp", "slp", "", "set large pages mode")
	rootCmd.PersistentFlags().BoolS("slt", "slt", false, "show technical information for l command")
	rootCmd.PersistentFlags().BoolS("snh", "snh", false, "store hard links as links")
	rootCmd.PersistentFlags().BoolS("sni", "sni", false, "store NT security information")
	rootCmd.PersistentFlags().BoolS("snl", "snl", false, "store symbolic links as links")
	rootCmd.PersistentFlags().StringS("sns", "sns", "", "store NTFS alternate streams")
	rootCmd.PersistentFlags().BoolS("so", "so", false, "write data to stdout")
	rootCmd.PersistentFlags().BoolS("spd", "spd", false, "disable wildcard matching for file names")
	rootCmd.PersistentFlags().BoolS("spe", "spe", false, "eliminate duplication of root folder for extract command")
	rootCmd.PersistentFlags().StringS("spf", "spf", "", "use fully qualified file paths")
	rootCmd.PersistentFlags().StringS("ssc", "ssc", "", "set sensitive case mode")
	rootCmd.PersistentFlags().BoolS("sse", "sse", false, "stop archive creating if input file cannot be opened")
	rootCmd.PersistentFlags().BoolS("ssp", "ssp", false, "do not change Last Access Time of source files")
	rootCmd.PersistentFlags().BoolS("ssw", "ssw", false, "compress shared files")
	rootCmd.PersistentFlags().BoolS("stl", "stl", false, "set archive timestamp from the most recently modified file")
	rootCmd.PersistentFlags().StringS("stm", "stm", "", "set CPU thread affinity mask")
	rootCmd.PersistentFlags().StringS("stx", "stx", "", "exclude archive type")
	rootCmd.PersistentFlags().StringS("t", "t", "", "set type of archive")
	rootCmd.PersistentFlags().StringS("u", "u", "", "update options")
	rootCmd.PersistentFlags().StringS("v", "v", "", "create volumes")
	rootCmd.PersistentFlags().StringS("w", "w", "", "set working directory")
	rootCmd.PersistentFlags().StringS("x", "x", "", "exclude filenames")
	rootCmd.PersistentFlags().BoolS("y", "y", false, "assume Yes on all queries")

	rootCmd.Flag("ai").NoOptDefVal = " "
	rootCmd.Flag("ai").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("ao").NoOptDefVal = " "
	rootCmd.Flag("ao").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("ax").NoOptDefVal = " "
	rootCmd.Flag("ax").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("bb").NoOptDefVal = " "
	rootCmd.Flag("bb").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("bs").NoOptDefVal = " "
	rootCmd.Flag("bs").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("i").NoOptDefVal = " "
	rootCmd.Flag("i").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("m").NoOptDefVal = " "
	rootCmd.Flag("m").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("o").NoOptDefVal = " "
	rootCmd.Flag("o").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("p").NoOptDefVal = " "
	rootCmd.Flag("p").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("r").NoOptDefVal = " "
	rootCmd.Flag("r").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("sa").NoOptDefVal = " "
	rootCmd.Flag("sa").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("scc").NoOptDefVal = " "
	rootCmd.Flag("scc").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("scrc").NoOptDefVal = " "
	rootCmd.Flag("scrc").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("scs").NoOptDefVal = " "
	rootCmd.Flag("scs").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("seml").NoOptDefVal = " "
	rootCmd.Flag("seml").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("sfx").NoOptDefVal = " "
	rootCmd.Flag("sfx").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("si").NoOptDefVal = " "
	rootCmd.Flag("si").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("slp").NoOptDefVal = " "
	rootCmd.Flag("slp").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("sns").NoOptDefVal = " "
	rootCmd.Flag("sns").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("spf").NoOptDefVal = " "
	rootCmd.Flag("spf").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("ssc").NoOptDefVal = " "
	rootCmd.Flag("ssc").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("stm").NoOptDefVal = " "
	rootCmd.Flag("stm").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("stx").NoOptDefVal = " "
	rootCmd.Flag("stx").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("t").NoOptDefVal = " "
	rootCmd.Flag("t").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("u").NoOptDefVal = " "
	rootCmd.Flag("u").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("v").NoOptDefVal = " "
	rootCmd.Flag("v").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("w").NoOptDefVal = " "
	rootCmd.Flag("w").OptargDelimiter = pflag.DelimiterDisabled
	rootCmd.Flag("x").NoOptDefVal = " "
	rootCmd.Flag("x").OptargDelimiter = pflag.DelimiterDisabled

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
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
