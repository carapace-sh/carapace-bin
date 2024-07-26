package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var passportPurgeCmd = &cobra.Command{
	Use:   "passport:purge [--revoked] [--expired] [--hours [HOURS]]",
	Short: "Purge revoked and / or expired tokens and authentication codes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(passportPurgeCmd).Standalone()

	passportPurgeCmd.Flags().Bool("expired", false, "Only purge expired tokens and authentication codes")
	passportPurgeCmd.Flags().String("hours", "", "The number of hours to retain expired tokens")
	passportPurgeCmd.Flags().Bool("revoked", false, "Only purge revoked tokens and authentication codes")
	rootCmd.AddCommand(passportPurgeCmd)
}
