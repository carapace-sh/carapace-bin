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
	teamCmd.PersistentFlags().Bool("json", false, "output as json")
	teamCmd.PersistentFlags().String("otp", "", "one-time password")
	teamCmd.PersistentFlags().BoolP("parseable", "p", false, "output parseable results")
	teamCmd.PersistentFlags().String("registry", "", "base URL of the npm registry")

	rootCmd.AddCommand(teamCmd)
}
