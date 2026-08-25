package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pi"
	"github.com/spf13/cobra"
)

var auth_print_api_keyCmd = &cobra.Command{
	Use:   "print-api-key",
	Short: "Print a provider API key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_print_api_keyCmd).Standalone()

	auth_print_api_keyCmd.Flags().String("model", "", "Model pattern or ID")
	auth_print_api_keyCmd.Flags().String("provider", "", "Provider name")
	authCmd.AddCommand(auth_print_api_keyCmd)

	carapace.Gen(auth_print_api_keyCmd).FlagCompletion(carapace.ActionMap{
		"model":    pi.ActionModels(),
		"provider": pi.ActionProviders(),
	})
}