package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:    "install",
	Short:  "Install the package",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolP("recursive", "R", false, "recurse directory")

	carapace.Gen(installCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if installCmd.Flag("recursive").Changed {
				return carapace.ActionDirectories()
			}
			return carapace.ActionFiles()
		}),
	)
}
