package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnosticsCmd = &cobra.Command{
	Use:   "diagnostics",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnosticsCmd).Standalone()

	diagnosticsCmd.Flags().Bool("disable-build-scripts", false, "Don't run build scripts or load `OUT_DIR` values by running `cargo check` before analysis.")
	diagnosticsCmd.Flags().Bool("disable-proc-macros", false, "Don't use expand proc macros.")
	rootCmd.AddCommand(diagnosticsCmd)

	carapace.Gen(diagnosticsCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
