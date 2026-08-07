package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var changelog_generateCmd = &cobra.Command{
	Use:   "generate [flags]",
	Short: "Generate a changelog for the current project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changelog_generateCmd).Standalone()

	changelog_generateCmd.Flags().String("config-file", "", "Path to the changelog configuration file in the project's Git repository. Defaults to '.gitlab/changelog_config.yml'.")
	changelog_generateCmd.Flags().String("date", "", "Date and time of the release, in ISO 8601 format (2016-03-11T03:45:40Z). Defaults to the current time.")
	changelog_generateCmd.Flags().String("from", "", "Start of the range of commits to use when generating the changelog, as a SHA. This commit is not included in the range.")
	changelog_generateCmd.Flags().String("to", "", "End of the range of commits to use when generating the changelog, as a SHA. This commit is included in the range. Defaults to the HEAD of the project's default branch.")
	changelog_generateCmd.Flags().String("trailer", "", "The Git trailer to use to include commits. Defaults to 'Changelog'.")
	changelog_generateCmd.Flags().StringP("version", "v", "", "Version to generate the changelog for. Must follow semantic versioning. Defaults to the version detected by 'git describe'.")
	changelogCmd.AddCommand(changelog_generateCmd)

	carapace.Gen(changelog_generateCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"from":        git.ActionRefs(git.RefOption{}.Default()),
		"to":          git.ActionRefs(git.RefOption{}.Default()),
	})
}
