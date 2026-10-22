package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var object_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save object locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_saveCmd).Standalone()

	object_saveCmd.Flags().String("file", "", "Destination filename (defaults to object name); using '-' as the filename will print the file to stdout")
	objectCmd.AddCommand(object_saveCmd)

	carapace.Gen(object_saveCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(object_saveCmd).PositionalCompletion(
		action.ActionContainers(),
		action.ActionObjects(),
	)
}
