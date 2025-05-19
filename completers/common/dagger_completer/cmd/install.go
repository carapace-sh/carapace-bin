package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dagger"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install [flags] MODULE",
	Short:   "Add a new dependency to a Dagger module",
	GroupID: "module",
	Aliases: []string{"use"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("focus", false, "Only show output for focused commands")
	installCmd.Flags().StringP("mod", "m", "", "Path to dagger.json config file for the module or a directory containing that file. Either local path (e.g. \"/path/to/some/dir\") or a github repo (e.g. \"github.com/dagger/dagger/path/to/some/subdir\")")
	installCmd.Flags().StringP("name", "n", "", "Name to use for the dependency in the module. Defaults to the name of the module being installed.")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"mod": dagger.ActionMods(),
	})

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return git.ActionRepositorySearch(git.SearchOpts{Github: true})
			default:
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: "https://" + c.Parts[0], Branches: true, Tags: true})
			}
		}),
	)
}
