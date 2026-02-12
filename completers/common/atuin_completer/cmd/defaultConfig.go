package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var defaultConfigCmd = &cobra.Command{
	Use:   "default-config",
	Short: "Print the default atuin configuration (config.toml)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(defaultConfigCmd).Standalone()

	defaultConfigCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(defaultConfigCmd)
}
