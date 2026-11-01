package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/paru_completer/cmd/common"
	"github.com/spf13/cobra"
)

var chrootctlCmd = &cobra.Command{
	Use:     "chrootctl",
	Aliases: []string{"C"},
	Short:   "Interactive shell to the chroot",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chrootctlCmd).Standalone()

	chrootctlCmd.Flags().BoolP("install", "i", false, "Install a package into the chroot")
	chrootctlCmd.Flags().BoolP("print", "p", false, "Print path to currently configured chroot")
	chrootctlCmd.Flags().BoolP("sysupgrade", "u", false, "Upgrade the chroot")
	common.AddNewFlags(chrootctlCmd)

	carapace.Gen(chrootctlCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if chrootctlCmd.Flag("install").Changed {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues()
		}),
	)
}
