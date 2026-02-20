package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lastUsedLayoutCmd = &cobra.Command{
	Use:   "last-used-layout",
	Short: "Switch to the last used layout",
}

func init() {
	lastUsedLayoutCmd.AddCommand(atCmd)
	carapace.Gen(lastUsedLayoutCmd).Standalone()

	lastUsedLayoutCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	lastUsedLayoutCmd.Flags().BoolP("all", "a", false, "Change the layout in all tabs")
	lastUsedLayoutCmd.Flags().StringP("match", "m", "", "The tab to match")
	lastUsedLayoutCmd.Flags().Bool("no-response", false, "Don't wait for a response from kitty")

	carapace.Gen(lastUsedLayoutCmd).FlagCompletion(carapace.ActionMap{})
}