package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/newrelic"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "newrelic",
	Short: "The New Relic CLI",
	Long:  "https://github.com/newrelic/newrelic-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().IntP("accountId", "a", 0, "the account ID to use. Can be overridden by setting NEW_RELIC_ACCOUNT_ID")
	rootCmd.PersistentFlags().Bool("debug", false, "debug level logging")
	rootCmd.PersistentFlags().String("format", "JSON", "output text format [JSON, Text, YAML]")
	rootCmd.Flags().BoolP("help", "h", false, "help for newrelic-dev")
	rootCmd.PersistentFlags().Bool("plain", false, "output compact text")
	rootCmd.PersistentFlags().String("profile", "", "the authentication profile to use")
	rootCmd.PersistentFlags().Bool("trace", false, "trace level logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"accountId": newrelic.ActionAccountIds(),
		"format":    carapace.ActionValues("JSON", "Text", "YAML"),
		"profile":   newrelic.ActionProfiles(),
	})
}
