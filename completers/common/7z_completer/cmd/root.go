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
	cmd.Flags().BoolS("an", "an", false, "disable archive_name field")
	cmd.Flags().StringS("ao", "ao", "", "set overwrite mode")
	cmd.Flags().StringS("bb", "bb", "", "set output log level")
	cmd.Flags().BoolS("bd", "bd", false, "disable progress indicator")
	cmd.Flags().StringS("bs", "bs", "", "set output stream for output/error/progress")
	cmd.Flags().BoolS("bt", "bt", false, "show execution time statistics")
	cmd.Flags().StringS("m", "m", "", "set compression method")
	cmd.Flags().StringS("o", "o", "", "set output directory")
	cmd.Flags().StringS("p", "p", "", "set password")
	cmd.Flags().BoolS("r", "r", false, "recurse subdirectories")
	cmd.Flags().StringS("sa", "sa", "", "set archive name mode")
	cmd.Flags().StringS("scc", "scc", "", "set charset for console input/output")
	cmd.Flags().StringS("scrc", "scrc", "", "set hash function")
	cmd.Flags().StringS("scs", "scs", "", "set charset for list files")
	cmd.Flags().BoolS("sdel", "sdel", false, "delete files after compression")
	cmd.Flags().BoolS("seml", "seml", false, "send archive by email")
	cmd.Flags().StringS("sfx", "sfx", "", "create SFX archive")
	cmd.Flags().StringS("si", "si", "", "read data from stdin")
	cmd.Flags().BoolS("slp", "slp", false, "set large pages mode")
	cmd.Flags().BoolS("slt", "slt", false, "show technical information for l command")
	cmd.Flags().BoolS("snh", "snh", false, "store hard links as links")
	cmd.Flags().BoolS("sni", "sni", false, "store NT security information")
	cmd.Flags().BoolS("snl", "snl", false, "store symbolic links as links")
	cmd.Flags().BoolS("sns", "sns", false, "store NTFS alternate streams")
	cmd.Flags().BoolS("so", "so", false, "write data to stdout")
	cmd.Flags().BoolS("spd", "spd", false, "disable wildcard matching for file names")
	cmd.Flags().BoolS("spe", "spe", false, "eliminate duplication of root folder for extract command")
	cmd.Flags().BoolS("spf", "spf", false, "use fully qualified file paths")
	cmd.Flags().BoolS("ssc", "ssc", false, "set sensitive case mode")
	cmd.Flags().BoolS("ssw", "ssw", false, "compress shared files")
	cmd.Flags().BoolS("stl", "stl", false, "set archive timestamp from the most recently modified file")
	cmd.Flags().StringS("stm", "stm", "", "set CPU thread affinity mask")
	cmd.Flags().StringS("stx", "stx", "", "exclude archive type")
	cmd.Flags().StringS("t", "t", "", "set type of archive")
	cmd.Flags().StringS("v", "v", "", "create volumes")
	cmd.Flags().StringS("w", "w", "", "set working directory")
	cmd.Flags().BoolS("y", "y", false, "assume Yes on all queries")

	cmd.Flag("ao").NoOptDefVal = " "
	cmd.Flag("bb").NoOptDefVal = " "
	cmd.Flag("bs").NoOptDefVal = " "
	cmd.Flag("m").NoOptDefVal = " "
	cmd.Flag("o").NoOptDefVal = " "
	cmd.Flag("p").NoOptDefVal = " "
	cmd.Flag("sa").NoOptDefVal = " "
	cmd.Flag("scc").NoOptDefVal = " "
	cmd.Flag("scrc").NoOptDefVal = " "
	cmd.Flag("scs").NoOptDefVal = " "
	cmd.Flag("sfx").NoOptDefVal = " "
	cmd.Flag("si").NoOptDefVal = " "
	cmd.Flag("stm").NoOptDefVal = " "
	cmd.Flag("stx").NoOptDefVal = " "
	cmd.Flag("t").NoOptDefVal = " "
	cmd.Flag("v").NoOptDefVal = " "
	cmd.Flag("w").NoOptDefVal = " "
}

// addCommonFlagCompletions registers flag completions for common switches.
func addCommonFlagCompletions(cmd *cobra.Command) {
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"ao":   carapace.ActionValues("a", "s", "t", "u"),
		"bb":   carapace.ActionValues("0", "1", "2", "3"),
		"o":    carapace.ActionDirectories(),
		"sa":   carapace.ActionValues("a", "e", "s"),
		"scc":  carapace.ActionValues("UTF-8", "WIN", "DOS"),
		"scrc": carapace.ActionValues("CRC32", "CRC64", "SHA1", "SHA256", "*"),
		"scs":  carapace.ActionValues("UTF-8", "UTF-16LE", "UTF-16BE", "WIN", "DOS"),
		"sfx":  carapace.ActionFiles(),
		"t": carapace.ActionValues(
			"7z", "bzip2", "cab", "gzip",
			"iso", "lzma", "lzma86",
			"tar", "wim", "xz", "zip",
		),
		"w": carapace.ActionDirectories(),
	})
}
