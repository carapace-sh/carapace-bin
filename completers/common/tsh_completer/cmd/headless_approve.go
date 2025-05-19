package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var headless_approveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve a headless authentication request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(headless_approveCmd).Standalone()

	headless_approveCmd.Flags().Bool("no-skip-confirm", false, "Skip confirmation and prompt for MFA immediately")
	headless_approveCmd.Flags().Bool("skip-confirm", false, "Skip confirmation and prompt for MFA immediately")
	headless_approveCmd.Flag("no-skip-confirm").Hidden = true
	headlessCmd.AddCommand(headless_approveCmd)
}
