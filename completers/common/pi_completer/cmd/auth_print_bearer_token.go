package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pi"
	"github.com/spf13/cobra"
)

var auth_print_bearer_tokenCmd = &cobra.Command{
	Use:   "print-bearer-token",
	Short: "Print an OAuth bearer token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_print_bearer_tokenCmd).Standalone()

	auth_print_bearer_tokenCmd.Flags().String("min-expiry", "", "Minimum expiry duration (e.g. 30m, 1h)")
	auth_print_bearer_tokenCmd.Flags().String("model", "", "Model pattern or ID")
	auth_print_bearer_tokenCmd.Flags().String("provider", "", "Provider name")
	authCmd.AddCommand(auth_print_bearer_tokenCmd)

	carapace.Gen(auth_print_bearer_tokenCmd).FlagCompletion(carapace.ActionMap{
		"provider": pi.ActionProviders(),
	})
}