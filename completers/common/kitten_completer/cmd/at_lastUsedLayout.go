package at

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

	carapace.Gen(lastUsedLayoutCmd).FlagCompletion(carapace.ActionMap{})
}