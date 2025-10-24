package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/carapace-sh/carapace/pkg/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add dependencies to a Cargo.toml manifest file",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("add"),
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().String("base", "", "The path base to use when adding from a local crate (unstable).")
	addCmd.Flags().String("branch", "", "Git branch to download the crate from")
	addCmd.Flags().Bool("build", false, "Add as build dependency")
	addCmd.Flags().Bool("default-features", false, "Re-enable the default features")
	addCmd.Flags().Bool("dev", false, "Add as development dependency")
	addCmd.Flags().BoolP("dry-run", "n", false, "Don't actually write the manifest")
	addCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	addCmd.Flags().String("git", "", "Git repository location")
	addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	addCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	addCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	addCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	addCmd.Flags().Bool("no-default-features", false, "Disable the default features")
	addCmd.Flags().Bool("no-optional", false, "Mark the dependency as required")
	addCmd.Flags().Bool("no-public", false, "Mark the dependency as private (unstable)")
	addCmd.Flags().Bool("optional", false, "Mark the dependency as optional")
	addCmd.Flags().StringP("package", "p", "", "Package to modify")
	addCmd.Flags().String("path", "", "Filesystem path to local crate to add")
	addCmd.Flags().Bool("public", false, "Mark the dependency as public (unstable)")
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
		"lockfile-path": carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(addCmd, true),
		"path":          carapace.ActionDirectories(),
	})

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionDirectories(),
			carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionGithubPackageSearch().NoSpace()
				case 1:
					return action.ActionGithubPackageVersions(c.Parts[0])
				default:
					return carapace.ActionValues()
				}
			}).UnlessF(condition.CompletingPathS),
		).ToA(),
	)
}
