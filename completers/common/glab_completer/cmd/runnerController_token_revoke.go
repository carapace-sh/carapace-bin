package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runnerController_token_revokeCmd = &cobra.Command{
	Use:   "revoke <controller-id> <token-id> [flags]",
	Short: "Revoke a token from a runner controller. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runnerController_token_revokeCmd).Standalone()

	runnerController_token_revokeCmd.Flags().BoolP("force", "f", false, "Skip confirmation prompt.")
	runnerController_tokenCmd.AddCommand(runnerController_token_revokeCmd)
}
