package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var plugin_import_from_sourcesCmd = &cobra.Command{
	Use:   "sources",
	Short: "Build a plugin from sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_import_from_sourcesCmd).Standalone()

	plugin_import_fromCmd.AddCommand(plugin_import_from_sourcesCmd)

	plugin_import_from_sourcesCmd.Flags().String("branch", "", "The branch of the repository that should be cloned")
	plugin_import_from_sourcesCmd.Flags().BoolP("force", "f", false, "Always clone the repository instead of trying to fetch the latest commits--json   Format the output as an NDJSON stream")
	plugin_import_from_sourcesCmd.Flags().Bool("no-minify", false, "Build a plugin for development (debugging) - non-minified and non-mangled")
	plugin_import_from_sourcesCmd.Flags().String("path", "", "The path where the repository should be cloned to")
	plugin_import_from_sourcesCmd.Flags().String("repository", "", "The repository that should be cloned")
	plugin_import_fromCmd.AddCommand(plugin_import_from_sourcesCmd)

	carapace.Gen(plugin_import_from_sourcesCmd).FlagCompletion(carapace.ActionMap{
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
