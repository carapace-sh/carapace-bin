package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:    "unpack",
	Short:  "Unpack the package",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpackCmd).Standalone()

	unpackCmd.Flags().BoolP("recursive", "R", false, "recurse directory")

	carapace.Gen(unpackCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if unpackCmd.Flag("recursive").Changed {
				return carapace.ActionDirectories()
			}
			return carapace.ActionFiles()
		}),
	)
}
