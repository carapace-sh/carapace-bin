package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systrayCmd = &cobra.Command{
	Use:   "systray",
	Short: "Run a systray application to manage Tailscale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systrayCmd).Standalone()

	systrayCmd.Flags().String("theme", "", "color theme for Tailscale icon")
	rootCmd.AddCommand(systrayCmd)

	carapace.Gen(systrayCmd).FlagCompletion(carapace.ActionMap{
		"theme": carapace.ActionValues("dark", "dark:nobg", "light", "light:nobg"),
	})
}
