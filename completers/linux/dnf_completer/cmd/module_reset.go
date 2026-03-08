package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleResetCmd = &cobra.Command{
	Use:   "reset <module-spec>...",
	Short: "reset module state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleResetCmd).Standalone()

	moduleResetCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	moduleResetCmd.Flags().Bool("skip-unavailable", false, "allow skipping modules")

	moduleCmd.AddCommand(moduleResetCmd)
}
