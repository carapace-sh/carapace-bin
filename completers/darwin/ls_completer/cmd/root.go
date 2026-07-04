package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ls",
	Short: "list directory contents",
	Long:  "https://keith.github.io/xcode-manpages/ls.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("1", "1", false, "Force output to be one entry per line")
	rootCmd.Flags().BoolS("@", "@", false, "Display extended attribute keys and sizes in long (-l) output")
	rootCmd.Flags().BoolS("A", "A", false, "List all entries except for . and ..")
	rootCmd.Flags().BoolS("B", "B", false, "Force printing of non-printable characters (as \\xxx) in file names")
	rootCmd.Flags().BoolS("C", "C", false, "Force multi-column output; this is the default when output is to a terminal")
	rootCmd.Flags().BoolS("F", "F", false, "Display a slash (`/') immediately after each pathname that is a directory")
	rootCmd.Flags().BoolS("G", "G", false, "Enable colorized output")
	rootCmd.Flags().BoolS("H", "H", false, "Symbolic links on the command line are followed")
	rootCmd.Flags().BoolS("I", "I", false, "Prevent -A, -a, -f from having an effect")
	rootCmd.Flags().BoolS("L", "L", false, "Follow all symbolic links")
	rootCmd.Flags().BoolS("O", "O", false, "Display the file flags in long (-l) output")
	rootCmd.Flags().BoolS("P", "P", false, "Do not follow symbolic links")
	rootCmd.Flags().BoolS("R", "R", false, "Recursively list subdirectories encountered")
	rootCmd.Flags().BoolS("S", "S", false, "Sort files by size")
	rootCmd.Flags().BoolS("T", "T", false, "When used with the -l (lowercase letter ``ell'') option, display complete time information")
	rootCmd.Flags().BoolS("U", "U", false, "Use time when file was created for sorting or printing")
	rootCmd.Flags().BoolS("W", "W", false, "Display whiteouts when scanning for the -A or -a options")
	rootCmd.Flags().BoolS("X", "X", false, "When sorting alphabetically, sort first by extension")
	rootCmd.Flags().BoolS("a", "a", false, "List all entries including those starting with a dot .")
	rootCmd.Flags().BoolS("b", "b", false, "Force printing of non-printable characters (as \\xxx) in file names")
	rootCmd.Flags().BoolS("c", "c", false, "Use time when file status was last changed for sorting or printing")
	rootCmd.Flags().BoolS("d", "d", false, "Directories are listed as plain files")
	rootCmd.Flags().BoolS("e", "e", false, "Print the Access Control List (ACL) associated with the file or directory")
	rootCmd.Flags().BoolS("f", "f", false, "Output is not sorted")
	rootCmd.Flags().BoolS("g", "g", false, "List the long entries, but omit the file owner")
	rootCmd.Flags().BoolS("h", "h", false, "Modifies the -s and -l options to use human readable numbers")
	rootCmd.Flags().BoolS("i", "i", false, "For each file, print the file's file serial number (inode number)")
	rootCmd.Flags().BoolS("k", "k", false, "Modifies the -s option to use 1024 byte units")
	rootCmd.Flags().BoolS("l", "l", false, "List in long format")
	rootCmd.Flags().BoolS("m", "m", false, "Stream output format")
	rootCmd.Flags().BoolS("n", "n", false, "List in long format, but omit the file owner")
	rootCmd.Flags().BoolS("o", "o", false, "List in long format, but omit the file group")
	rootCmd.Flags().BoolS("p", "p", false, "Write a slash (`/') after each filename if that file is a directory")
	rootCmd.Flags().BoolS("q", "q", false, "Force printing of non-graphic characters in file names as the character `?'")
	rootCmd.Flags().BoolS("r", "r", false, "Reverse the order of the sort to get reverse lexicographical order")
	rootCmd.Flags().BoolS("s", "s", false, "Display the number of file system blocks actually used by each file")
	rootCmd.Flags().BoolS("t", "t", false, "Sort by time modified (most recently modified first)")
	rootCmd.Flags().BoolS("u", "u", false, "Use time of last access, instead of last modification of the file for sorting")
	rootCmd.Flags().BoolS("v", "v", false, "Force unedited output, even if the output would use non-graphic characters")
	rootCmd.Flags().BoolS("w", "w", false, "Force raw printing of non-printable characters")
	rootCmd.Flags().BoolS("x", "x", false, "The same as -C, except that the multi-column output is produced with entries sorted across")
	rootCmd.Flags().BoolS("y", "y", false, "When used with -l, the file flags are printed")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
