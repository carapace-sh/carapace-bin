package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup zellij and check its configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().Bool("check", false, "Checks the configuration of zellij and displays currently used directories")
	setupCmd.Flags().Bool("clean", false, "Disables loading of configuration file at default location, loads the defaults that zellij ships with")
	setupCmd.Flags().Bool("dump-config", false, "Dump the default configuration file to stdout")
	setupCmd.Flags().String("dump-layout", "", "Dump specified layout to stdout")
	setupCmd.Flags().String("dump-plugins", "", "Dump the builtin plugins to DIR or \"DATA DIR\" if unspecified")
	setupCmd.Flags().String("dump-swap-layout", "", "Dump the specified swap layout file to stdout")
	setupCmd.Flags().String("generate-auto-start", "", "Generates auto-start script for the specified shell")
	setupCmd.Flags().String("generate-completion", "", "Generates completion for the specified shell")
	setupCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(setupCmd)

	carapace.Gen(setupCmd).FlagCompletion(carapace.ActionMap{
		"dump-plugins":        carapace.ActionFiles(),
		"generate-auto-start": carapace.ActionValues("bash", "zsh", "fish", "powershell", "elvish"),
		"generate-completion": carapace.ActionValues("bash", "zsh", "fish", "powershell", "elvish"),
	})
}
