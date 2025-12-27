package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download [pattern]...",
	Short: "download package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downloadCmd).Standalone()
	rootCmd.AddCommand(downloadCmd)

	carapace.Gen(downloadCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch(),
	)
}
