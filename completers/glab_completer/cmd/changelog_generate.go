package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var changelog_generateCmd = &cobra.Command{
	Use:   "generate [flags]",
	Short: "Generate a changelog for the repository/project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changelog_generateCmd).Standalone()

	changelog_generateCmd.Flags().String("config-file", "", "The path of changelog configuration file in the project's Git repository. Defaults to .gitlab/changelog_config.yml.")
	changelog_generateCmd.Flags().String("date", "", "The date and time of the release. Uses ISO 8601 (2016-03-11T03:45:40Z) format. Defaults to the current time.")
	changelog_generateCmd.Flags().String("from", "", "The start of the range of commits (as a SHA) to use for generating the changelog. This commit itself isn't included in the list.")
	changelog_generateCmd.Flags().String("to", "", "The end of the range of commits (as a SHA) to use for the changelog. This commit is included in the list. Defaults to the HEAD of the default project branch.")
	changelog_generateCmd.Flags().String("trailer", "", "The Git trailer to use for including commits. Defaults to Changelog.")
	changelog_generateCmd.Flags().StringP("version", "v", "", "The version to generate the changelog for. The format must follow semantic versioning. Defaults to the version of the local checkout (like using git describe).")
	changelogCmd.AddCommand(changelog_generateCmd)

	carapace.Gen(changelog_generateCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"from":        git.ActionRefs(git.RefOption{}.Default()),
		"to":          git.ActionRefs(git.RefOption{}.Default()),
	})
}
