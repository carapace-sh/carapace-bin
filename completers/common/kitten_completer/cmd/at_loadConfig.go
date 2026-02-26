package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_loadConfigCmd = &cobra.Command{
	Use:   "load-config",
	Short: "(Re)load a config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_loadConfigCmd).Standalone()

	at_loadConfigCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_loadConfigCmd.Flags().Bool("ignore-overrides", false, "Have previous config overrides ignored")
	at_loadConfigCmd.Flags().Bool("no-response", false, "Don't wait for a response indicating the success of the action")
	at_loadConfigCmd.Flags().StringP("override", "o", "", "Override individual configuration options")
	atCmd.AddCommand(at_loadConfigCmd)

}
