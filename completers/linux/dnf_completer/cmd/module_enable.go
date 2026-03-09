package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleEnableCmd = &cobra.Command{
	Use:   "enable <module-spec>...",
	Short: "enable module streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleEnableCmd).Standalone()

	moduleEnableCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	moduleEnableCmd.Flags().Bool("skip-unavailable", false, "allow skipping modules")

	moduleCmd.AddCommand(moduleEnableCmd)
}
