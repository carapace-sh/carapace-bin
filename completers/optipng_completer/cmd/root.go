package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "optipng",
	Short: "Optimize Portable Network Graphics files",
	Long:  "http://optipng.sourceforge.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("backup", false, "keep a backup of the modified files")
	rootCmd.Flags().Bool("clobber", false, "overwrite existing files")
	rootCmd.Flags().String("dir", "", "write output file(s) to <directory>")
	rootCmd.Flags().String("f", "", "PNG delta filters (0-5)")
	rootCmd.Flags().Bool("fix", false, "enable error recovery")
	rootCmd.Flags().Bool("force", false, "enforce writing of a new output file")
	rootCmd.Flags().Bool("full", false, "produce a full report on IDAT (might reduce speed)")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().String("i", "", "PNG interlace type (0-1)")
	rootCmd.Flags().Bool("keep", false, "keep a backup of the modified files")
	rootCmd.Flags().String("log", "", "log messages to <file>")
	rootCmd.Flags().Bool("nb", false, "no bit depth reduction")
	rootCmd.Flags().Bool("nc", false, "no color type reduction")
	rootCmd.Flags().Bool("np", false, "no palette reduction")
	rootCmd.Flags().Bool("nx", false, "no reductions")
	rootCmd.Flags().Bool("nz", false, "no IDAT recoding")
	rootCmd.Flags().String("o", "", "optimization level (0-7)")
	rootCmd.Flags().String("out", "", "write output file to <file>")
	rootCmd.Flags().Bool("preserve", false, "preserve file attributes if possible")
	rootCmd.Flags().Bool("quiet", false, "run in quiet mode")
	rootCmd.Flags().Bool("silent", false, "run in quiet mode")
	rootCmd.Flags().Bool("simulate", false, "run in simulation mode")
	rootCmd.Flags().Bool("snip", false, "cut one image out of multi-image or animation files")
	rootCmd.Flags().String("strip", "", "strip metadata objects (e.g. \"all\")")
	rootCmd.Flags().String("zc", "", "zlib compression levels (1-9)")
	rootCmd.Flags().String("zm", "", "zlib memory levels (1-9)")
	rootCmd.Flags().String("zs", "", "zlib compression strategies (0-3)")
	rootCmd.Flags().String("zw", "", "zlib window size (256,512,1k,2k,4k,8k,16k,32k)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
		"f": carapace.ActionValuesDescribed(
			"0", "None",
			"1", "Left",
			"2", "Up",
			"3", "Average",
			"4", "Paeth",
			"5", "Adaptive",
		),
		"i": carapace.ActionValuesDescribed(
			"0", "non-interlaced",
			"1", "interlaced",
		),
		"log": carapace.ActionFiles(),
		"o":   carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9"),
		"out": carapace.ActionFiles(),
		"zc": carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
			if len(c.Parts) < 2 {
				return carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7")
			}
			return carapace.ActionValues()
		}),
		"zm": carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
			if len(c.Parts) < 2 {
				return carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9")
			}
			return carapace.ActionValues()
		}),
		"zs": carapace.ActionMultiParts("-", func(c carapace.Context) carapace.Action {
			if len(c.Parts) < 2 {
				return carapace.ActionValues("1", "2", "3")
			}
			return carapace.ActionValues()
		}),
		"zw": carapace.ActionValues("32k", "16k", "8k", "4k", "2k", "1k", "512", "256"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
