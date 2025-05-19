package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Opens the glab configuration file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_editCmd).Standalone()

	config_editCmd.Flags().BoolP("local", "l", false, "Open '.git/glab-cli/config.yml' file instead of the global '~/.config/glab-cli/config.yml' file.")
	configCmd.AddCommand(config_editCmd)
}
