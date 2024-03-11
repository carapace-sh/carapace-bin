package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var set_version_from_sourcesCmd = &cobra.Command{
	Use:   "sources",
	Short: "Build Yarn from master",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_version_from_sourcesCmd).Standalone()

	set_version_from_sourcesCmd.Flags().String("branch", "", "The branch of the repository that should be cloned")
	set_version_from_sourcesCmd.Flags().BoolP("force", "f", false, "Always clone the repository instead of trying to fetch the latest commits")
	set_version_from_sourcesCmd.Flags().Bool("no-minify", false, "Build a bundle for development (debugging) - non-minified and non-mangled")
	set_version_from_sourcesCmd.Flags().String("path", "", "The path where the repository should be cloned to")
	set_version_from_sourcesCmd.Flags().String("plugin", "", "An array of additional plugins that should be included in the bundle")
	set_version_from_sourcesCmd.Flags().String("repository", "", "The repository that should be cloned")
	set_version_from_sourcesCmd.Flags().Bool("skip-plugins", false, "Skip updating the contrib plugins")
	set_version_fromCmd.AddCommand(set_version_from_sourcesCmd)

	// TODO plugin
	carapace.Gen(set_version_from_sourcesCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := plugin_import_from_sourcesCmd.Flag("repository"); flag.Changed {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: flag.Value.String(), Branches: true, Tags: true})
			}
			return carapace.ActionValues()
		}),
		"path":       carapace.ActionDirectories(),
		"repository": git.ActionRepositorySearch(git.SearchOpts{}.Default()),
	})
}
