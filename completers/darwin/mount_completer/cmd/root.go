package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mount"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mount",
	Short: "mount file systems",
	Long:  "https://keith.github.io/xcode-man-pages/mount.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "mount all filesystems mentioned in fstab")
	rootCmd.Flags().BoolP("dry-run", "d", false, "dry run; skip the mount(2) syscall")
	rootCmd.Flags().BoolP("force", "f", false, "force revocation of write access when downgrading to read-only")
	rootCmd.Flags().BoolP("fskit", "F", false, "force filesystem type to be considered as an FSModule")
	rootCmd.Flags().BoolP("nofollow", "k", false, "don't follow symlinks in mount-on directory")
	rootCmd.Flags().StringP("options", "o", "", "comma-separated list of mount options")
	rootCmd.Flags().BoolP("read-only", "r", false, "mount the filesystem read-only")
	rootCmd.Flags().BoolP("read-write", "w", false, "mount the filesystem read-write")
	rootCmd.Flags().StringP("types", "t", "", "limit the set of filesystem types")
	rootCmd.Flags().BoolP("update", "u", false, "change the status of an already mounted filesystem")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"options": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			keys := make([]string, 0)
			for _, part := range c.Parts {
				keys = append(keys, strings.Split(part, "=")[0])
			}
			return mount.ActionMountOptions().Invoke(c).Filter(keys...).ToA().NoSpace()
		}),
		"types": fs.ActionFilesystemTypes().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}
			return mount.ActionSources()
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}
			if c.Args[0] != "" {
				return carapace.ActionDirectories()
			}
			return carapace.ActionValues()
		}),
	)
}
