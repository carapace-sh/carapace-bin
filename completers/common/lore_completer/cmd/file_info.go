package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var file_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about the given file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_infoCmd).Standalone()

	file_infoCmd.Flags().Bool("filtered", false, "If given, calculate the repository size based on the current local filter")
	file_infoCmd.Flags().BoolP("help", "h", false, "Print help")
	file_infoCmd.Flags().Bool("local", false, "If given, calculate the local file system size and hash based on the current local filter")
	file_infoCmd.Flags().String("revision", "", "Revision to get info from")
	file_infoCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	fileCmd.AddCommand(file_infoCmd)

	carapace.Gen(file_infoCmd).FlagCompletion(carapace.ActionMap{
		"revision": action.ActionRevisions(file_infoCmd),
		"targets":  carapace.ActionFiles(),
	})

	carapace.Gen(file_infoCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
