package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apt_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var sourceCmd = &cobra.Command{
	Use:   "source [pattern]...",
	Short: "download source package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sourceCmd).Standalone()

	sourceCmd.Flags().BoolP("build", "b", false, "compile source packages after downloading them")
	sourceCmd.Flags().BoolP("build-profiles", "P", false, "activated build profiles")
	sourceCmd.Flags().Bool("compile", false, "compile source packages after downloading them")
	sourceCmd.Flags().Bool("debian-only", false, "download only the debian file")
	sourceCmd.Flags().Bool("diff-only", false, "download only the diff")
	sourceCmd.Flags().Bool("dry-run", false, "perform a simulation of events taken")
	sourceCmd.Flags().Bool("dsc-only", false, "download only the dsc file")
	sourceCmd.Flags().BoolP("simulate", "s", false, "perform a simulation of events taken")
	sourceCmd.Flags().Bool("tar-only", false, "download only the tar file")
	// sourceCmd.Flags().BoolP("target-release", "t", false, "target release")
	common.AddGetFlags(sourceCmd)
	rootCmd.AddCommand(sourceCmd)

	carapace.Gen(sourceCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch().FilterArgs(),
	)
}
