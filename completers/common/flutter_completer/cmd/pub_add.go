package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pub"
	"github.com/spf13/cobra"
)

var pub_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a dependency to pubspec.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_addCmd).Standalone()

	pub_addCmd.Flags().StringP("directory", "C", "", "Run this in the directory <dir>.")
	pub_addCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change.")
	pub_addCmd.Flags().Bool("example", false, "Also update dependencies in `example/` after modifying pubspec.yaml in the root package (if it exists).")
	pub_addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_addCmd.Flags().Bool("no-example", false, "Do not update dependencies in `example/`.")
	pub_addCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	pub_addCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	pubCmd.AddCommand(pub_addCmd)

	carapace.Gen(pub_addCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
	})

	carapace.Gen(pub_addCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return pub.ActionPackageSearch().Invoke(c).Suffix(":").ToA()
			case 1:
				return pub.ActionPackageVersions(c.Parts[0])
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
