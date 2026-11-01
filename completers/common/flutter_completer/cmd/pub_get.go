package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get packages in a Flutter project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_getCmd).Standalone()

	pub_getCmd.Flags().StringP("directory", "C", "", "Run this in the directory <dir>.")
	pub_getCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but do not change any.")
	pub_getCmd.Flags().Bool("enforce-lockfile", false, "Enforce pubspec.lock. Fail `pub get` if the current `pubspec.lock` does not exactly specify a valid resolution of `pubspec.yaml` or if any content hash of a hosted package has changed.")
	pub_getCmd.Flags().Bool("example", false, "Also run in `example/` (if it exists).")
	pub_getCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_getCmd.Flags().Bool("no-example", false, "Do not run in `example/`.")
	pub_getCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	pub_getCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	pubCmd.AddCommand(pub_getCmd)

	carapace.Gen(pub_getCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
