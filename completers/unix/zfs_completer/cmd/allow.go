package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var allowCmd = &cobra.Command{
	Use:     "allow [-dglu] user|group[,...] perm|@setname[,...] filesystem|volume",
	Short:   "delegate permissions on a dataset",
	GroupID: "delegation",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(allowCmd).Standalone()

	allowCmd.Flags().BoolS("c", "c", false, "set create-time permissions")
	allowCmd.Flags().BoolS("d", "d", false, "allow for descendants only")
	allowCmd.Flags().BoolS("e", "e", false, "delegate to everyone")
	allowCmd.Flags().BoolS("g", "g", false, "delegate to group")
	allowCmd.Flags().BoolS("l", "l", false, "allow locally only")
	allowCmd.Flags().StringS("s", "s", "", "define or add to permission set")
	allowCmd.Flags().BoolS("u", "u", false, "delegate to user")

	rootCmd.AddCommand(allowCmd)

	carapace.Gen(allowCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if allowCmd.Flag("c").Changed || allowCmd.Flag("e").Changed || allowCmd.Flag("s").Changed {
				switch len(c.Args) {
				case 0:
					return zfs.ActionPermissions().UniqueList(",")
				default:
					return zfs.ActionFilesystems()
				}
			}
			switch len(c.Args) {
			case 0:
				switch {
				case allowCmd.Flag("u").Changed && allowCmd.Flag("g").Changed:
					return carapace.Batch(os.ActionUsers(), os.ActionGroups()).ToA().UniqueList(",")
				case allowCmd.Flag("u").Changed:
					return os.ActionUsers().UniqueList(",")
				case allowCmd.Flag("g").Changed:
					return os.ActionGroups().UniqueList(",")
				default:
					return carapace.Batch(
						carapace.ActionValues("everyone"),
						os.ActionUsers(),
						os.ActionGroups(),
						zfs.ActionFilesystems(),
					).ToA()
				}
			case 1:
				return zfs.ActionPermissions().UniqueList(",")
			default:
				return zfs.ActionFilesystems()
			}
		}),
	)
}
