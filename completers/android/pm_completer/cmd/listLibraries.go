package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listLibrariesCmd = &cobra.Command{
	Use:   "libraries",
	Short: "Print all system libraries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listLibrariesCmd).Standalone()

	listLibrariesCmd.Flags().BoolP("verbose", "v", false, "shows the location of the library in the device's filesystem")

	listCmd.AddCommand(listLibrariesCmd)

}
