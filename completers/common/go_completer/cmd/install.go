package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "compile and install packages and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().SetInterspersed(false)

	installCmd.Flags().BoolS("i", "i", false, "install the packages that are dependencies of the target")
	installCmd.Flags().BoolS("json", "json", false, "Emit build output in JSON suitable for automated processing")
	installCmd.Flags().StringS("o", "o", "", "set output file or directory")
	addBuildFlags(installCmd)
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionFiles(),
	})

	carapace.Gen(installCmd).PositionalCompletion(
		golang.ActionModuleSearch(),
	)
}
