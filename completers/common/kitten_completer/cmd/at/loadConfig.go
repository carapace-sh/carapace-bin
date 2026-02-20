package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loadConfigCmd = &cobra.Command{
	Use:   "load-config",
	Short: "(Re)load a config file",
}

func init() {
	loadConfigCmd.AddCommand(atCmd)
	carapace.Gen(loadConfigCmd).Standalone()

	loadConfigCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(loadConfigCmd).FlagCompletion(carapace.ActionMap{})
}