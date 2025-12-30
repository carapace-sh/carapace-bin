package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version <pkgname...>",
	Short: "Prints version of installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCommand).Standalone()

	rootCmd.AddCommand(versionCommand)

	carapace.Gen(versionCommand).PositionalCompletion(xbps.ActionPackages())
}
