package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var buildDepCmd = &cobra.Command{
	Use:   "build-dep",
	Short: "build-dep causes apt-get to install/remove packages in an attempt to satisfy the build dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildDepCmd).Standalone()

	rootCmd.AddCommand(buildDepCmd)

	carapace.Gen(buildDepCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch().FilterArgs(),
	)
}
