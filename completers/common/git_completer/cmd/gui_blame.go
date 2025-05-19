package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var gui_blameCmd = &cobra.Command{
	Use:   "blame",
	Short: "Start a blame viewer on the specified file on the given version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gui_blameCmd).Standalone()
	gui_blameCmd.Flags().String("line", "", "Loads annotations as described above and automatically scrolls the view")

	guiCmd.AddCommand(gui_blameCmd)

	carapace.Gen(gui_blameCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			git.ActionRefs(git.RefOption{}.Default()).UnlessF(condition.CompletingPath),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRefFiles(c.Args[0]).UnlessF(condition.File(c.Args[0]))
		}),
	)
}
