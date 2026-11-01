package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleEnableCmd = &cobra.Command{
	Use:   "enable [options] <module-spec>...",
	Short: "enable module streams and make their packages available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleEnableCmd).Standalone()

	moduleEnableCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	moduleEnableCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")

	moduleCmd.AddCommand(moduleEnableCmd)
}
