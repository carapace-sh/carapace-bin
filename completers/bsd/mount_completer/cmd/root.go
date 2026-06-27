package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mount"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mount",
	Short: "mount file systems",
	Long:  "https://manpage.me/?mount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "All the filesystems listed via getfsent are mounted")
	rootCmd.Flags().BoolS("d", "d", false, "Causes everything to be done except for the actual system call")
	rootCmd.Flags().BoolS("f", "f", false, "Forces the revocation of write access")
	rootCmd.Flags().BoolS("F", "F", false, "Forces the file system type be considered as an FSModule")
	rootCmd.Flags().BoolS("k", "k", false, "Do not follow any symlinks that may be present")
	rootCmd.Flags().StringS("o", "o", "", "A comma separated string of options")
	rootCmd.Flags().BoolS("r", "r", false, "Mount the file system read-only")
	rootCmd.Flags().StringS("t", "t", "", "Indicate the file system type")
	rootCmd.Flags().BoolS("u", "u", false, "Update the status of an already mounted file system")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")
	rootCmd.Flags().BoolS("w", "w", false, "Mount the file system read-write")

	rootCmd.Flags().SetInterspersed(false)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"o": mount.ActionMountOptions().UniqueList(","),
		"t": fs.ActionFilesystemTypes().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("a").Changed {
				return carapace.ActionValues()
			}
			return carapace.Batch(
				fs.ActionBlockDevices(),
				carapace.ActionFiles(),
			).ToA()
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
