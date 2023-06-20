package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace/pkg/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add dependency to a Cargo.toml manifest file",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("add"),
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().String("branch", "", "Git branch to download the crate from")
	addCmd.Flags().Bool("build", false, "Add as build dependency")
	addCmd.Flags().Bool("default-features", false, "Re-enable the default features")
	addCmd.Flags().Bool("dev", false, "Add as development dependency")
	addCmd.Flags().Bool("dry-run", false, "Don't actually write the manifest")
	addCmd.Flags().StringSliceP("features", "F", []string{}, "Space or comma separated list of features to activate")
	addCmd.Flags().String("git", "", "Git repository location")
	addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	addCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	addCmd.Flags().Bool("no-default-features", false, "Disable the default features")
	addCmd.Flags().Bool("no-optional", false, "Mark the dependency as required")
	addCmd.Flags().Bool("optional", false, "Mark the dependency as optional")
	addCmd.Flags().StringP("package", "p", "", "Package to modify")
	addCmd.Flags().String("path", "", "Filesystem path to local crate to add")
	addCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	addCmd.Flags().String("registry", "", "Package registry for this dependency")
	addCmd.Flags().String("rename", "", "Rename the dependency")
	addCmd.Flags().String("rev", "", "Git reference to download the crate from")
	addCmd.Flags().String("tag", "", "Git tag to download the crate from")
	addCmd.Flags().String("target", "", "Add as dependency to the given target platform")
	rootCmd.AddCommand(addCmd)

	// TODO flag completion
	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := addCmd.Flag("git"); f.Changed && strings.HasPrefix(f.Value.String(), "https://github.com/") {
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: f.Value.String(), Branches: true, Tags: true})
			}
			return carapace.ActionValues()
		}),
		"features": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 1 && !util.HasPathPrefix(c.Args[0]) {
				// TODO limit to specific version
				return action.ActionGithubPackageFeatures(strings.Split(c.Args[0], "@")[0]).UniqueList(",")
			}
			return carapace.ActionValues()
		}),
		"git": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRepositorySearch(git.SearchOpts{}.Default())
		}),
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(addCmd, true),
		"path":          carapace.ActionDirectories(),
	})

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.Value) {
				return carapace.ActionDirectories()
			}

			return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionGithubPackageSearch().NoSpace()
				case 1:
					return action.ActionGithubPackageVersions(c.Parts[0])
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
