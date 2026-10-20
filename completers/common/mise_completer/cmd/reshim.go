package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var reshimCmd = &cobra.Command{
	Use:   "reshim",
	Short: "Create shims for executables provided by installed tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reshimCmd).Standalone()

	reshimCmd.Flags().BoolP("force", "f", false, "Rebuild all mise-owned shims")
	reshimCmd.Flags().BoolP("help", "h", false, "Print help")
	reshimCmd.Flags().Bool("system", false, "Rebuild the system shims, for tools installed with `mise install --system`")
	rootCmd.AddCommand(reshimCmd)

	carapace.Gen(reshimCmd).PositionalCompletion(
		mise.ActionTools(),
	)
}
