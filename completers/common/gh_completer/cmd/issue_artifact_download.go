package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_artifact_downloadCmd = &cobra.Command{
	Use:   "download {<issue> [<artifact-number>...] | <artifact-url>...}",
	Short: "Download issue artifacts (preview)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_artifact_downloadCmd).Standalone()

	issue_artifact_downloadCmd.Flags().Bool("clobber", false, "Overwrite existing files of the same name")
	issue_artifact_downloadCmd.Flags().StringP("dir", "D", ".", "The `directory` to download files into")
	issue_artifact_downloadCmd.Flags().StringP("output", "O", "", "The `file` to write a single artifact to (use \"-\" to write to standard output)")
	issue_artifact_downloadCmd.Flags().Bool("skip-existing", false, "Skip downloading when files of the same name exist")
	issue_artifact_downloadCmd.Flags().String("version", "", "Download the version with this `number` from the edit history")
	issue_artifactCmd.AddCommand(issue_artifact_downloadCmd)

	carapace.Gen(issue_artifact_downloadCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
		"output": carapace.Batch(
			carapace.ActionValuesDescribed("-", "standard output"),
			carapace.ActionFiles(),
		).ToA(),
		"version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 1 {
				return action.ActionIssueArtifactVersions(issue_artifact_downloadCmd, c.Args[0], c.Args[1])
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(issue_artifact_downloadCmd).PositionalCompletion(
		action.ActionIssues(issue_artifact_downloadCmd, action.IssueOpts{Open: true, Closed: true}),
	)

	carapace.Gen(issue_artifact_downloadCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionIssueArtifacts(issue_artifact_downloadCmd, c.Args[0])
		}),
	)
}
