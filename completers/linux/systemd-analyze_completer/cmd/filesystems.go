package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var filesystemsCmd = &cobra.Command{
	Use:   "filesystems",
	Short: "List known filesystems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filesystemsCmd).Standalone()

	rootCmd.AddCommand(filesystemsCmd)

	carapace.Gen(filesystemsCmd).PositionalAnyCompletion(
		action.ActionFilesystemSets(),
	)
}
