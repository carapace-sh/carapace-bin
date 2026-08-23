package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pax",
	Short: "read and write file archives and copy directory hierarchies",
	Long:  "https://man.freebsd.org/cgi/man.cgi?pax",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "Use NUL as pathname terminator")
	rootCmd.Flags().StringS("B", "B", "", "Number of bytes written to a single volume")
	rootCmd.Flags().StringS("E", "E", "", "Limit for read errors")
	rootCmd.Flags().StringS("G", "G", "", "Select files based on group")
	rootCmd.Flags().BoolS("H", "H", false, "Follow only command line symbolic links")
	rootCmd.Flags().BoolS("L", "L", false, "Follow all symbolic links")
	rootCmd.Flags().BoolS("O", "O", false, "Force the archive to be one volume")
	rootCmd.Flags().BoolS("P", "P", false, "Do not follow symbolic links")
	rootCmd.Flags().StringS("T", "T", "", "Select files based on modification time")
	rootCmd.Flags().StringS("U", "U", "", "Select files based on user")
	rootCmd.Flags().BoolS("X", "X", false, "Do not traverse past mount points")
	rootCmd.Flags().BoolS("Y", "Y", false, "Same as -D")
	rootCmd.Flags().BoolS("Z", "Z", false, "Same as -u")
	rootCmd.Flags().BoolS("a", "a", false, "Append files to the end of an archive")
	rootCmd.Flags().StringS("b", "b", "", "Blocksize")
	rootCmd.Flags().BoolS("c", "c", false, "Match all files except those specified")
	rootCmd.Flags().BoolS("d", "d", false, "Cause files of type directory to be archived")
	rootCmd.Flags().StringS("f", "f", "", "Archive file")
	rootCmd.Flags().BoolS("i", "i", false, "Interactively rename files")
	rootCmd.Flags().BoolS("j", "j", false, "Use bzip2 compression")
	rootCmd.Flags().BoolS("k", "k", false, "Do not overwrite existing files")
	rootCmd.Flags().BoolS("l", "l", false, "Link files")
	rootCmd.Flags().BoolS("n", "n", false, "Select the first archive member that matches")
	rootCmd.Flags().StringS("o", "o", "", "Archive format options")
	rootCmd.Flags().StringS("p", "p", "", "File mode and ownership options")
	rootCmd.Flags().BoolS("r", "r", false, "Read an archive file")
	rootCmd.Flags().StringS("s", "s", "", "Modify file or archive member names")
	rootCmd.Flags().BoolS("t", "t", false, "Reset access times")
	rootCmd.Flags().BoolS("u", "u", false, "Ignore files that are older")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
	rootCmd.Flags().BoolS("w", "w", false, "Write an archive file")
	rootCmd.Flags().StringS("x", "x", "", "Archive format")
	rootCmd.Flags().BoolS("z", "z", false, "Use gzip compression")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
		"x": carapace.ActionValues("cpio", "sv4cpio", "sv4crc", "tar", "ustar", "pax"),
	})
}
