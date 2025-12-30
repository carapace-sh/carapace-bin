package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getpkgdepnameCommand = &cobra.Command{
	Use:   "getpkgdepname <dep...>",
	Short: "Prints pkgname from dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getpkgdepnameCommand).Standalone()

	rootCmd.AddCommand(getpkgdepnameCommand)
}
