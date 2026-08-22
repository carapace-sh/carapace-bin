package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update apps, or Scoop itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()
	updateCmd.Flags().BoolP("all", "a", false, "update all apps (alternative to `*`)")
	updateCmd.Flags().BoolP("force", "f", false, "force update even when there isn't a newer version")
	updateCmd.Flags().BoolP("global", "g", false, "update a globally installed app")
	updateCmd.Flags().BoolP("independent", "i", false, "don't install dependencies automatically")
	updateCmd.Flags().BoolP("no-cache", "k", false, "don't use the download cache")
	updateCmd.Flags().BoolP("quiet", "q", false, "hide extraneous messages")
	updateCmd.Flags().BoolP("skip-hash-check", "s", false, "skip hash validation")
	rootCmd.AddCommand(updateCmd)
}
