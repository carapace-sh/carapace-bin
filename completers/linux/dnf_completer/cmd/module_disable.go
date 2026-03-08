package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleDisableCmd = &cobra.Command{
	Use:   "disable <module-spec>...",
	Short: "disable modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleDisableCmd).Standalone()

	moduleDisableCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	moduleDisableCmd.Flags().Bool("skip-unavailable", false, "allow skipping modules")

	moduleCmd.AddCommand(moduleDisableCmd)
}
