package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Manage organization teams and team memberships",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teamCmd).Standalone()

	teamCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	teamCmd.Flags().Bool("json", false, "Output results as JSON")
	teamCmd.Flags().String("otp", "", "One-time password for registries that require two-factor authentication")
	teamCmd.Flags().Bool("parseable", false, "Output parseable results (tab-separated)")
	teamCmd.Flags().String("registry", "", "The base URL of the npm registry")
	rootCmd.AddCommand(teamCmd)
}
