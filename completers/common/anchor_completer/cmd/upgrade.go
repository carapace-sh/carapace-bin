package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:    "upgrade",
	Short:  "Deploys, initializes an IDL, and migrates all in one command. Upgrades a single program. The configured wallet must be the upgrade authority",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().BoolP("help", "h", false, "Print help")
	upgradeCmd.Flags().String("max-retries", "0", "Max times to retry on failure")
	upgradeCmd.Flags().StringP("program-id", "p", "", "The program to upgrade")
	upgradeCmd.MarkFlagRequired("program-id")
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
