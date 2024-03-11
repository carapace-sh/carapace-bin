package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alpha_watchCmd = &cobra.Command{
	Use:   "watch [SERVICE...]",
	Short: "Watch build context for service and rebuild/refresh containers when files are updated",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_watchCmd).Standalone()

	alpha_watchCmd.Flags().Bool("no-up", false, "Do not build & start services before watching")
	alpha_watchCmd.Flags().Bool("quiet", false, "hide build output")
	alphaCmd.AddCommand(alpha_watchCmd)
}
