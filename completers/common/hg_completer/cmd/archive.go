package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:     "archive",
	Short:   "create an unversioned archive of a repository revision",
	GroupID: groups[group_change_import_export].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(archiveCmd).Standalone()

	archiveCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	archiveCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	archiveCmd.Flags().Bool("no-decode", false, "do not pass files through decoders")
	archiveCmd.Flags().StringP("prefix", "p", "", "directory prefix for files in archive")
	archiveCmd.Flags().StringP("rev", "r", "", "revision to distribute")
	archiveCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	archiveCmd.Flags().StringP("type", "t", "", "type of distribution to create")
	rootCmd.AddCommand(archiveCmd)

	carapace.Gen(archiveCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"rev":     action.ActionRevisions(),
		"type":    carapace.ActionValues("files", "tar", "tgz", "txz", "uzip", "zip"),
	})

	carapace.Gen(archiveCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
