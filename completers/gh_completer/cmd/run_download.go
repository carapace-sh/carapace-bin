package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download artifacts generated by a workflow run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_downloadCmd).Standalone()
	run_downloadCmd.Flags().StringP("dir", "D", ".", "The directory to download artifacts into")
	run_downloadCmd.Flags().StringArrayP("name", "n", []string{}, "Only download artifacts that match any of the given names")
	runCmd.AddCommand(run_downloadCmd)

	carapace.Gen(run_downloadCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
		// TODO name
	})

	carapace.Gen(run_downloadCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
		"name": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionWorkflowArtifactNames(run_downloadCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(run_downloadCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_downloadCmd, action.RunOpts{Successful: true}),
	)
}
