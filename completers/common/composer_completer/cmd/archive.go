package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Creates an archive of this composer package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(archiveCmd).Standalone()

	archiveCmd.Flags().String("dir", "", "Write the archive to this directory")
	archiveCmd.Flags().String("file", "", "Write the archive with the given file name. Note that the format will be appended.")
	archiveCmd.Flags().StringP("format", "f", "", "Format of the resulting archive: tar, tar.gz, tar.bz2 or zip (default tar)")
	archiveCmd.Flags().Bool("ignore-filters", false, "Ignore filters when saving package")

	rootCmd.AddCommand(archiveCmd)

	carapace.Gen(archiveCmd).FlagCompletion(carapace.ActionMap{
		"dir":    carapace.ActionDirectories(),
		"format": carapace.ActionValues("tar", "tar.bz2", "tar.gz", "zip"),
	})
}
