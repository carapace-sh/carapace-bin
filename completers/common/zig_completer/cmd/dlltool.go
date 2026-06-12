package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dlltoolCmd = &cobra.Command{
	Use:                "dlltool",
	Short:              "Use Zig as a drop-in dlltool.exe",
	Run:                func(cmd *cobra.Command, args []string) {},
	GroupID:            "wrapper",
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(dlltoolCmd).Standalone()

	rootCmd.AddCommand(dlltoolCmd)

	carapace.Gen(dlltoolCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("dlltool"),
	)
}
