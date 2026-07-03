package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tree",
	Short: "list contents of directories in a tree-like format",
	Long:  "https://linux.die.net/man/1/tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Print ANSI lines graphic indentation lines.")
	rootCmd.Flags().BoolS("C", "C", false, "Turn colorization on always.")
	rootCmd.Flags().BoolS("D", "D", false, "Print the date of last modification or (-c) status change.")
	rootCmd.Flags().BoolS("F", "F", false, "Appends '/', '=', '*', '@', '|' or '>' as per ls -F.")
	rootCmd.Flags().StringS("H", "H", "", "Prints out HTML format with baseHREF as top directory.")
	rootCmd.Flags().StringS("I", "I", "", "Do not list files that match the given pattern.")
	rootCmd.Flags().BoolS("J", "J", false, "Prints out an JSON representation of the tree.")
	rootCmd.Flags().StringS("L", "L", "", "Descend only level directories deep.")
	rootCmd.Flags().BoolS("N", "N", false, "Print non-printable characters as is.")
	rootCmd.Flags().StringS("P", "P", "", "List only those files that match the pattern given.")
	rootCmd.Flags().BoolS("Q", "Q", false, "Quote filenames with double quotes.")
	rootCmd.Flags().BoolS("R", "R", false, "Rerun tree when max dir level reached.")
	rootCmd.Flags().BoolS("S", "S", false, "Print with CP437 (console) graphics indentation lines.")
	rootCmd.Flags().StringS("T", "T", "", "Replace the default HTML title and H1 header with string.")
	rootCmd.Flags().BoolS("U", "U", false, "Leave files unsorted.")
	rootCmd.Flags().BoolS("X", "X", false, "Prints out an XML representation of the tree.")
	rootCmd.Flags().BoolS("a", "a", false, "All files are listed.")
	rootCmd.Flags().Bool("acl", false, "Print permissions with a + if an ACL is present.")
	rootCmd.Flags().String("authority", "", "Set OSC 8 hyperlink authority/hostname.")
	rootCmd.Flags().BoolS("c", "c", false, "Sort files by last status change time.")
	rootCmd.Flags().String("charset", "", "Use charset X for terminal/HTML and indentation line output.")
	rootCmd.Flags().String("compress", "", "Compress indentation lines.")
	rootCmd.Flags().Bool("condense", false, "Condense directory singletons to a single line of output.")
	rootCmd.Flags().BoolS("d", "d", false, "List directories only.")
	rootCmd.Flags().Bool("device", false, "Print device ID number to which each file belongs.")
	rootCmd.Flags().Bool("dirsfirst", false, "List directories before files (-U disables).")
	rootCmd.Flags().Bool("du", false, "Compute size of directories by their contents.")
	rootCmd.Flags().BoolS("f", "f", false, "Print the full path prefix for each file.")
	rootCmd.Flags().Bool("fflinks", false, "Process link information when using --fromfile.")
	rootCmd.Flags().String("filelimit", "", "Do not descend dirs with more than # files in them.")
	rootCmd.Flags().Bool("filesfirst", false, "List files before directories (-U disables).")
	rootCmd.Flags().Bool("fromfile", false, "Reads paths from files (.=stdin)")
	rootCmd.Flags().Bool("fromtabfile", false, "Reads trees from tab indented files (.=stdin)")
	rootCmd.Flags().BoolS("g", "g", false, "Displays file group owner or GID number.")
	rootCmd.Flags().String("gitfile", "", "Explicitly read a gitignore file.")
	rootCmd.Flags().Bool("gitignore", false, "Filter by using .gitignore files.")
	rootCmd.Flags().BoolS("h", "h", false, "Print the size in a more human readable way.")
	rootCmd.Flags().Bool("help", false, "Print usage and this help message and exit.")
	rootCmd.Flags().String("hintro", "", "Use file X as the HTML intro.")
	rootCmd.Flags().String("houtro", "", "Use file X as the HTML outro.")
	rootCmd.Flags().Bool("hyperlink", false, "Turn on OSC 8 terminal hyperlinks.")
	rootCmd.Flags().BoolS("i", "i", false, "Don't print indentation lines.")
	rootCmd.Flags().Bool("ignore-case", false, "Ignore case when pattern matching.")
	rootCmd.Flags().Bool("info", false, "Print information about files found in .info files.")
	rootCmd.Flags().String("infofile", "", "Explicitly read info file.")
	rootCmd.Flags().Bool("inodes", false, "Print inode number of each file.")
	rootCmd.Flags().BoolS("l", "l", false, "Follow symbolic links like directories.")
	rootCmd.Flags().Bool("matchdirs", false, "Include directory names in -P pattern matching.")
	rootCmd.Flags().Bool("metafirst", false, "Print meta-data at the beginning of each line.")
	rootCmd.Flags().BoolS("n", "n", false, "Turn colorization off always (-C overrides).")
	rootCmd.Flags().Bool("nolinks", false, "Turn off hyperlinks in HTML output.")
	rootCmd.Flags().Bool("noreport", false, "Turn off file/directory count at end of tree listing.")
	rootCmd.Flags().StringS("o", "o", "", "Output to file instead of stdout.")
	rootCmd.Flags().Bool("opt-toggle", false, "Enable option toggling.")
	rootCmd.Flags().BoolS("p", "p", false, "Print the protections for each file.")
	rootCmd.Flags().Bool("prune", false, "Prune empty directories from the output.")
	rootCmd.Flags().BoolS("q", "q", false, "Print non-printable characters as '?'.")
	rootCmd.Flags().BoolS("r", "r", false, "Reverse the order of the sort.")
	rootCmd.Flags().BoolS("s", "s", false, "Print the size in bytes of each file.")
	rootCmd.Flags().String("scheme", "", "Set OSC 8 hyperlink scheme, default file://")
	rootCmd.Flags().Bool("selinux", false, "Print the selinux security label if present.")
	rootCmd.Flags().Bool("si", false, "Like -h, but use in SI units (powers of 1000).")
	rootCmd.Flags().String("sort", "", "Select sort: name,version,size,mtime,ctime,none.")
	rootCmd.Flags().BoolS("t", "t", false, "Sort files by last modification time.")
	rootCmd.Flags().String("timefmt", "", "Print and format time according to the format fmt.")
	rootCmd.Flags().BoolS("u", "u", false, "Displays file owner or UID number.")
	rootCmd.Flags().BoolS("v", "v", false, "Sort files alphanumerically by version.")
	rootCmd.Flags().Bool("version", false, "Print version and exit.")
	rootCmd.Flags().BoolS("x", "x", false, "Stay on current filesystem only.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"gitfile":  carapace.ActionFiles(),
		"hintro":   carapace.ActionFiles(),
		"houtro":   carapace.ActionFiles(),
		"infofile": carapace.ActionFiles(),
		"o":        carapace.ActionFiles(),
		"sort":     carapace.ActionValues("name", "version", "size", "mtime", "ctime", "none"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
