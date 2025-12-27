package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/paru_completer/cmd/common"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"B"},
	Short:   "Build PKGBUILDs on disk",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().BoolP("install", "i", false, "Install package as well as building")
	common.AddNewFlags(buildCmd)

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
