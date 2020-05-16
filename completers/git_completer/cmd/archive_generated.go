package cmd

import (
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Create an archive of files from a named tree",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	archiveCmd.Flags().BoolP("0", "0", false, "store only")
	archiveCmd.Flags().BoolP("1", "1", false, "compress faster")
	archiveCmd.Flags().BoolP("9", "9", false, "compress better")
	archiveCmd.Flags().String("exec", "", "path to the remote git-upload-archive command")
	archiveCmd.Flags().String("format", "", "archive format")
	archiveCmd.Flags().BoolP("list", "l", false, "list supported archive formats")
	archiveCmd.Flags().BoolP("output", "o", false, "<file>   write the archive to this file")
	archiveCmd.Flags().String("prefix", "", "prepend prefix to each pathname in the archive")
	archiveCmd.Flags().String("remote", "", "retrieve the archive from remote repository <repo>")
	archiveCmd.Flags().BoolP("verbose", "v", false, "report archived files on stderr")
	archiveCmd.Flags().Bool("worktree-attributes", false, "read .gitattributes in working directory")
	rootCmd.AddCommand(archiveCmd)
}
