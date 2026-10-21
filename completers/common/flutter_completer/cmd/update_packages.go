package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var update_packagesCmd = &cobra.Command{
	Use:     "update-packages",
	Short:   "Update the packages inside the Flutter repo. This is intended for CI and repo maintainers. Normal Flutter developers should not have to use this command",
	Aliases: []string{"upgrade-packages"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(update_packagesCmd).Standalone()

	update_packagesCmd.Flags().StringSlice("cherry-pick", nil, "Attempt to update only the specified package. To be specified as [pub package name]:[pub package version],[pub package2 name]:[pub package2 version].")
	update_packagesCmd.Flags().Bool("crash", false, "For Flutter CLI testing only, forces this command to throw an unhandled exception.")
	update_packagesCmd.Flags().Bool("exclude-tools", false, "Don't update the deps in tools. For example when unpinning a dep.")
	update_packagesCmd.Flags().Bool("force-upgrade", false, "Attempt to update all the dependencies to their latest versions. This will actually modify the pubspec.yaml files in your checkout.")
	update_packagesCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	update_packagesCmd.Flags().Bool("no-exclude-tools", false, "Don't update the deps in tools. For example when unpinning a dep.")
	update_packagesCmd.Flags().Bool("no-upgrade-major", false, "Upgrade major versions as well. Only makes sense with force-upgrade.")
	update_packagesCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	update_packagesCmd.Flags().Bool("update-hashes", false, "Update the hashes of the pubspecs.")
	update_packagesCmd.Flags().Bool("upgrade-major", false, "Upgrade major versions as well. Only makes sense with force-upgrade.")
	rootCmd.AddCommand(update_packagesCmd)
}
