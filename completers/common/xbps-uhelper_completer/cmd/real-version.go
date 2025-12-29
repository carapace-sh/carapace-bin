package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var realversionCommand = &cobra.Command{
	Use:   "real-version <pkgname...>",
	Short: "Prints version of installed real packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(realversionCommand).Standalone()

	rootCmd.AddCommand(realversionCommand)

	carapace.Gen(realversionCommand).PositionalCompletion(xbps.ActionPackages())
}
