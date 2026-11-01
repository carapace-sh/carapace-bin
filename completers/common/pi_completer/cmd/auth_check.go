package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pi"
	"github.com/spf13/cobra"
)

var auth_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check provider readiness",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_checkCmd).Standalone()

	auth_checkCmd.Flags().Bool("credentials", false, "Emit the credential (or include it in JSON)")
	auth_checkCmd.Flags().Bool("json", false, "Output result as JSON")
	auth_checkCmd.Flags().String("model", "", "Model pattern or ID")
	auth_checkCmd.Flags().Bool("no-refresh", false, "Do not refresh expired OAuth credentials")
	auth_checkCmd.Flags().String("provider", "", "Provider name")
	authCmd.AddCommand(auth_checkCmd)

	carapace.Gen(auth_checkCmd).FlagCompletion(carapace.ActionMap{
		"model":    pi.ActionModels(),
		"provider": pi.ActionProviders(),
	})
}
