package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "7z",
	Short: "command-line port of 7zip",
	Long:  "https://linux.die.net/man/1/7z",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolS("ai", "ai", false, "include archives")
	rootCmd.Flags().BoolS("an", "an", false, "disable archive_name field")
	rootCmd.Flags().BoolS("ao", "ao", false, "set overwrite mode")
	rootCmd.Flags().BoolS("ax", "ax", false, "exclude archives")
	rootCmd.Flags().BoolS("bd", "bd", false, "disable progress indicator")
	rootCmd.Flags().BoolS("bs", "bs", false, "set output stream for output/error/progress line")
	rootCmd.Flags().BoolS("bt", "bt", false, "show execution time statistics")
	rootCmd.Flags().BoolS("i", "i", false, "include filenames")
	rootCmd.Flags().BoolS("m", "m", false, "set compression method")
	rootCmd.Flags().BoolS("o", "o", false, "set output directory")
	rootCmd.Flags().BoolS("p", "p", false, "set password")
	rootCmd.Flags().BoolS("r", "r", false, "recurse subdirectories")
	rootCmd.Flags().BoolS("sa", "sa", false, "set archive name mode")
	rootCmd.Flags().BoolS("scc", "scc", false, "set charset for console input/output")
	rootCmd.Flags().BoolS("scs", "scs", false, "set charset for list files")
	rootCmd.Flags().BoolS("sdel", "sdel", false, "delete files after compression")
	rootCmd.Flags().BoolS("seml", "seml", false, "send archive by email")
	rootCmd.Flags().BoolS("sfx", "sfx", false, "create self-extracting archive")
	rootCmd.Flags().BoolS("si", "si", false, "read data from stdin")
	rootCmd.Flags().BoolS("slp", "slp", false, "set large pages mode")
	rootCmd.Flags().BoolS("slt", "slt", false, "show technical information")
	rootCmd.Flags().BoolS("snh", "snh", false, "store hard links as links")
	rootCmd.Flags().BoolS("snl", "snl", false, "store symbolic links as links")
	rootCmd.Flags().BoolS("sni", "sni", false, "store NT security information")
	rootCmd.Flags().BoolS("sns", "sns", false, "store NTFS alternate streams")
	rootCmd.Flags().BoolS("so", "so", false, "write data to stdout")
	rootCmd.Flags().BoolS("spd", "spd", false, "disable wildcard matching")
	rootCmd.Flags().BoolS("spe", "spe", false, "eliminate duplication of root folder")
	rootCmd.Flags().BoolS("spf", "spf", false, "use fully qualified file paths")
	rootCmd.Flags().BoolS("ssc", "ssc", false, "set sensitive case mode")
	rootCmd.Flags().BoolS("ssw", "ssw", false, "compress shared files")
	rootCmd.Flags().BoolS("stl", "stl", false, "set archive timestamp from newest file")
	rootCmd.Flags().BoolS("t", "t", false, "set archive type")
	rootCmd.Flags().BoolS("u", "u", false, "update options")
	rootCmd.Flags().BoolS("v", "v", false, "create volumes")
	rootCmd.Flags().BoolS("w", "w", false, "set working directory")
	rootCmd.Flags().BoolS("x", "x", false, "exclude filenames")
	rootCmd.Flags().BoolS("y", "y", false, "assume yes on all queries")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionDirectories(),
		"t": carapace.ActionValues("7z", "bzip2", "gzip", "tar", "wim", "xz", "zip"),
		"w": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"a", "add files to archive",
			"b", "benchmark",
			"d", "delete files from archive",
			"e", "extract files from archive without paths",
			"h", "calculate hash values",
			"i", "show information about supported formats",
			"l", "list contents of archive",
			"rn", "rename files in archive",
			"t", "test integrity of archive",
			"u", "update files to archive",
			"x", "extract files with full paths",
		),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) < 2 {
				return archiveFiles()
			}

			switch c.Args[0] {
			case "d", "e", "l", "rn", "t", "x":
				return fs.ActionZipFileContents(c.Args[1]).Invoke(c).Filter(c.Args[2:]...).ToA()
			default:
				return carapace.ActionFiles().FilterArgs()
			}
		}),
	)
}

func archiveFiles() carapace.Action {
	return carapace.ActionFiles(
		".7z", ".apk", ".apm", ".ar", ".arj", ".bz2", ".bzip2", ".cab", ".chm",
		".cpio", ".cramfs", ".deb", ".dmg", ".esd", ".fat", ".gz", ".gzip",
		".hfs", ".iso", ".jar", ".lha", ".lzh", ".lzma", ".msi", ".nsis",
		".ntfs", ".rar", ".rpm", ".squashfs", ".swm", ".tar", ".taz", ".tbz",
		".tbz2", ".tgz", ".txz", ".udf", ".vhd", ".wim", ".xar", ".xz",
		".z", ".zip",
	)
}
