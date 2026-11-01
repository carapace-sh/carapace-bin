package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap___applyAccountPlanCmd = &cobra.Command{
	Use:    "__apply-account-plan",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap___applyAccountPlanCmd).Standalone()

	bootstrap___applyAccountPlanCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap___applyAccountPlanCmd)
}
