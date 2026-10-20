// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap___applySystemPlanCmd = &cobra.Command{
	Use:    "__apply-system-plan",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap___applySystemPlanCmd).Standalone()

	bootstrap___applySystemPlanCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap___applySystemPlanCmd)
}
