package cmd

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
	loadConfigCmd.Flags().Bool("ignore-overrides", false, "Have previous config overrides ignored")
	loadConfigCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	loadConfigCmd.Flags().StringP("override", "o", "", "Override individual configuration options")

	carapace.Gen(loadConfigCmd).FlagCompletion(carapace.ActionMap{})
}