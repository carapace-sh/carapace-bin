package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var realVersionCommand = &cobra.Command{
	Use:   "real-version <pkgname...>",
	Short: "Prints version of installed real packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(realVersionCommand).Standalone()

	rootCmd.AddCommand(realVersionCommand)

	carapace.Gen(realVersionCommand).PositionalCompletion(xbps.ActionPackages())
}
