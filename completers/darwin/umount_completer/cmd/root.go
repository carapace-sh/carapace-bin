package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mount"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umount",
	Short: "unmount filesystems",
	Long:  "https://keith.github.io/xcode-man-pages/umount.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "unmount all filesystems mentioned in fstab")
	rootCmd.Flags().BoolP("all-mounted", "A", false, "unmount all currently mounted filesystems except the root")
	rootCmd.Flags().BoolP("force", "f", false, "forcibly unmount the filesystem")
	rootCmd.Flags().StringP("host", "h", "", "only unmount filesystems mounted from the specified host")
	rootCmd.Flags().StringP("types", "t", "", "limit the set of filesystem types")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"types": fs.ActionFilesystemTypes().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("all").Changed ||
				rootCmd.Flag("all-mounted").Changed {
				return carapace.ActionValues()
			}
			return carapace.Batch(
				mount.ActionSources(),
				fs.ActionMounts(),
			).ToA()
		}),
	)
}
