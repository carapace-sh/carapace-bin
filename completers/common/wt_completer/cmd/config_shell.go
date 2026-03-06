package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Shell integration setup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_shellCmd).Standalone()

	config_shellCmd.Flags().BoolP("help", "h", false, "Print help")
	configCmd.AddCommand(config_shellCmd)
}
