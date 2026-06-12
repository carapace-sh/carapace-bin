package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var libCmd = &cobra.Command{
	Use:     "lib",
	Short:   "Use Zig as a drop-in lib.exe",
	GroupID: "wrapper",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(libCmd).Standalone()

	rootCmd.AddCommand(libCmd)

	carapace.Gen(libCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("lib"),
	)
}
