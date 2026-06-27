package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umount",
	Short: "Unmount filesystems",
	Long:  "https://manpage.me/?umount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "All the filesystems described via getfsent are unmounted")
	rootCmd.Flags().BoolS("A", "A", false, "All the currently mounted filesystems except the root are unmounted")
	rootCmd.Flags().BoolS("f", "f", false, "The filesystem is forcibly unmounted")
	rootCmd.Flags().StringS("h", "h", "", "Only filesystems mounted from the specified host will be unmounted")
	rootCmd.Flags().StringS("t", "t", "", "Indicate the actions should only be taken on filesystems of the specified type")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose, additional information is printed out as each filesystem is unmounted")

	rootCmd.Flags().SetInterspersed(false)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"t": fs.ActionFilesystemTypes().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("a").Changed || rootCmd.Flag("A").Changed {
				return carapace.ActionValues()
			}
			return fs.ActionMounts()
		}),
	)
}
