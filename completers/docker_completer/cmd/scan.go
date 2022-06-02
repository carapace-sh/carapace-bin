package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:                "scan",
	Short:              "A tool to scan your images",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(scanCmd).Standalone()

	rootCmd.AddCommand(scanCmd)

	carapace.Gen(scanCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("docker-scan"),
	)
}
