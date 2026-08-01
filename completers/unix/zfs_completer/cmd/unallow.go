package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var unallowCmd = &cobra.Command{
	Use:     "unallow [-dglru] user|group[,...] [perm|@setname[,...]] filesystem|volume",
	Short:   "revoke delegated permissions",
	GroupID: "delegation",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unallowCmd).Standalone()

	unallowCmd.Flags().BoolS("c", "c", false, "remove create-time permissions")
	unallowCmd.Flags().BoolS("d", "d", false, "remove from descendants only")
	unallowCmd.Flags().BoolS("e", "e", false, "remove from everyone")
	unallowCmd.Flags().BoolS("g", "g", false, "remove from group")
	unallowCmd.Flags().BoolS("l", "l", false, "remove locally only")
	unallowCmd.Flags().BoolS("r", "r", false, "recursively remove")
	unallowCmd.Flags().StringS("s", "s", "", "remove from permission set")
	unallowCmd.Flags().BoolS("u", "u", false, "remove from user")

	rootCmd.AddCommand(unallowCmd)

	carapace.Gen(unallowCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if unallowCmd.Flag("c").Changed || unallowCmd.Flag("e").Changed || unallowCmd.Flag("s").Changed {
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
				case unallowCmd.Flag("u").Changed && unallowCmd.Flag("g").Changed:
					return carapace.Batch(os.ActionUsers(), os.ActionGroups()).ToA().UniqueList(",")
				case unallowCmd.Flag("u").Changed:
					return os.ActionUsers().UniqueList(",")
				case unallowCmd.Flag("g").Changed:
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
				return carapace.Batch(
					zfs.ActionPermissions().UniqueList(","),
					zfs.ActionFilesystems(),
				).ToA()
			default:
				return zfs.ActionFilesystems()
			}
		}),
	)
}
