package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeNixCmd = &cobra.Command{
	Use:     "upgrade-nix",
	Short:   "upgrade Nix to the stable version declared in Nixpkgs",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeNixCmd).Standalone()

	upgradeNixCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	upgradeNixCmd.Flags().String("nix-store-paths-url", "", "The URL of the file that contains the store paths of the latest Nix release")
	upgradeNixCmd.Flags().StringP("profile", "p", "", "The path to the Nix profile to upgrade")
	rootCmd.AddCommand(upgradeNixCmd)

	addLoggingFlags(upgradeNixCmd)

	carapace.Gen(upgradeNixCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionFiles(),
	})
}
