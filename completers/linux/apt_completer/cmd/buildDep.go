package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apt_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var buildDepCmd = &cobra.Command{
	Use:   "build-dep package...",
	Short: "install build dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildDepCmd).Standalone()

	buildDepCmd.Flags().BoolP("build-profiles", "P", false, "activated build profiles")
	buildDepCmd.Flags().Bool("dry-run", false, "perform a simulation of events taken")
	buildDepCmd.Flags().BoolP("host-architecture", "a", false, "architecture")
	buildDepCmd.Flags().Bool("no-strict-pinning", false, "consider all versions of a package")
	buildDepCmd.Flags().Bool("purge", false, "remove and purge packages")
	buildDepCmd.Flags().BoolP("simulate", "s", false, "perform a simulation of events taken")
	buildDepCmd.Flags().String("solver", "", "set solver")
	buildDepCmd.Flags().StringP("target-release", "t", "", "target-release")
	common.AddGetFlags(buildDepCmd)
	rootCmd.AddCommand(buildDepCmd)

	carapace.Gen(buildDepCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch().FilterArgs(),
	)
}
