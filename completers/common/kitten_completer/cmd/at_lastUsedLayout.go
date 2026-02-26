package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_lastUsedLayoutCmd = &cobra.Command{
	Use:   "last-used-layout",
	Short: "Switch to the last used layout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_lastUsedLayoutCmd).Standalone()

	at_lastUsedLayoutCmd.Flags().BoolP("all", "a", false, "Change the layout in all tabs")
	at_lastUsedLayoutCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_lastUsedLayoutCmd.Flags().StringP("match", "m", "", "The tab to match")
	at_lastUsedLayoutCmd.Flags().Bool("no-response", false, "Don't wait for a response from kitty")
	atCmd.AddCommand(at_lastUsedLayoutCmd)

}
