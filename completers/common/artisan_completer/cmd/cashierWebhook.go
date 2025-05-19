package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cashierWebhookCmd = &cobra.Command{
	Use:   "cashier:webhook [--disabled] [--url [URL]] [--api-version [API-VERSION]]",
	Short: "Create the Stripe webhook to interact with Cashier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cashierWebhookCmd).Standalone()

	cashierWebhookCmd.Flags().String("api-version", "", "The Stripe API version the webhook should use")
	cashierWebhookCmd.Flags().Bool("disabled", false, "Immediately disable the webhook after creation")
	cashierWebhookCmd.Flags().String("url", "", "The URL endpoint for the webhook")
	rootCmd.AddCommand(cashierWebhookCmd)
}
