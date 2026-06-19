package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/common"
	"github.com/spf13/cobra"
)

var developCmd = &cobra.Command{
	Use:   "develop",
	Short: "Install the crate as module in the current virtualenv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(developCmd).Standalone()

	developCmd.Flags().StringP("bindings", "b", "", "Which kind of bindings to use")
	developCmd.Flags().Bool("compression-enable-large-file-support", false, "Whether to use large file support for ZIP files")
	developCmd.Flags().String("compression-level", "", "Zip compression level. Defaults to method default")
	developCmd.Flags().String("compression-method", "deflated", "Zip compression method")
	developCmd.Flags().StringP("extras", "E", "", "Install extra requires aka. optional dependencies")
	developCmd.Flags().Bool("generate-stubs", false, "Auto generate Python type stubs by introspecting the binary")
	developCmd.Flags().StringP("group", "G", "", "Install dependency groups defined in pyproject.toml (PEP 735)")
	developCmd.Flags().String("pip-path", "", "Use a specific pip installation instead of the default one")
	developCmd.Flags().BoolP("release", "r", false, "Pass --release to cargo")
	developCmd.Flags().Bool("skip-install", false, "Skip installation, only build the extension module inplace")
	developCmd.Flags().Bool("strip", false, "Strip the library for minimum file size")
	developCmd.Flags().Bool("uv", false, "Use `uv` to install packages instead of `pip`")
	common.AddCargoFlags(developCmd)
	rootCmd.AddCommand(developCmd)

	carapace.Gen(developCmd).FlagCompletion(carapace.ActionMap{
		"bindings":           action.ActionBindings(),
		"compression-method": action.ActionCompressionMethods(),
		"pip-path":           carapace.ActionFiles(),
	})
}
