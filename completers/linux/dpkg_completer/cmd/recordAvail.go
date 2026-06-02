package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recordAvailCmd = &cobra.Command{
	Use:    "record-avail",
	Short:  "Update which packages are available with information from the package",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recordAvailCmd).Standalone()

	recordAvailCmd.Flags().BoolP("recursive", "R", false, "recurse directory")

	carapace.Gen(recordAvailCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if recordAvailCmd.Flag("recursive").Changed {
				return carapace.ActionDirectories()
			}
			return carapace.ActionFiles()
		}),
	)
}
