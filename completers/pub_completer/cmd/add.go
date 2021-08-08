package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pub_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/git"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a dependency to pubspec.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolP("dev", "d", false, "Adds package to the development dependencies instead.")
	addCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change any.")
	addCmd.Flags().String("git-path", "", "Path of git package in repository")
	addCmd.Flags().String("git-ref", "", "Git branch or commit to be retrieved")
	addCmd.Flags().String("git-url", "", "Git URL of the package")
	addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	addCmd.Flags().String("hosted-url", "", "URL of package host server")
	addCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	addCmd.Flags().Bool("no-precompile", false, "Do not precompile executables in immediate dependencies.")
	addCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	addCmd.Flags().String("path", "", "Local path")
	addCmd.Flags().Bool("precompile", false, "Precompile executables in immediate dependencies.")
	addCmd.Flags().String("sdk", "", "SDK source for package")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		// TODO "git-path":
		"git-ref": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := addCmd.Flag("git-url"); flag.Changed {
				return git.ActionLsRemoteRefs(flag.Value.String(), git.LsRemoteRefOption{Branches: true, Tags: true})
			}
			return carapace.ActionMessage("git-url not set")
		}),
		"path": carapace.ActionDirectories(),
		// TODO "sdk":
	})

	carapace.Gen(addCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionPackageSearch().Invoke(c).Suffix(":").ToA()
			case 1:
				return action.ActionPackageVersions(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
