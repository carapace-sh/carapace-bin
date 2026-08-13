package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var teams_ssoCmd = &cobra.Command{
	Use:   "sso",
	Short: "Show SAML/SSO configuration for the current team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teams_ssoCmd).Standalone()

	teams_ssoCmd.Flags().String("format", "", "Output format")
	teams_ssoCmd.Flags().Bool("json", false, "Output as JSON")

	teamsCmd.AddCommand(teams_ssoCmd)

	carapace.Gen(teams_ssoCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
