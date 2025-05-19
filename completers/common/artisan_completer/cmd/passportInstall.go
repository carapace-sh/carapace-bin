package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var passportInstallCmd = &cobra.Command{
	Use:   "passport:install [--uuids] [--force] [--length [LENGTH]]",
	Short: "Run the commands necessary to prepare Passport for use",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(passportInstallCmd).Standalone()

	passportInstallCmd.Flags().Bool("force", false, "Overwrite keys they already exist")
	passportInstallCmd.Flags().String("length", "", "The length of the private key")
	passportInstallCmd.Flags().Bool("uuids", false, "Use UUIDs for all client IDs")
	rootCmd.AddCommand(passportInstallCmd)
}
