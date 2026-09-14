package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use:     "bundle",
	Short:   "bundle an application so that it works outside of the Nix store",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundleCmd).Standalone()

	bundleCmd.Flags().String("bundler", "", "Use a custom bundler instead of the default")
	bundleCmd.Flags().StringP("out-link", "o", "", "Override the name of the symlink to the build result")
	rootCmd.AddCommand(bundleCmd)

	common.AddEvaluationFlags(bundleCmd)
	common.AddFlakeFlags(bundleCmd)
	common.AddInterpretationFlags(bundleCmd)
	common.AddLoggingFlags(bundleCmd)

	carapace.Gen(bundleCmd).FlagCompletion(carapace.ActionMap{
		"bundler": carapace.Batch(
			carapace.ActionDirectories(),
			nix.ActionFlakes(),
		).ToA(),
	})
	carapace.Gen(bundleCmd).PositionalCompletion(nix.ActionInstallables())
}
