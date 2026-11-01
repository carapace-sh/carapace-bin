package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_macos_defaults_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the state of the configured macOS preferences",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_macos_defaults_statusCmd).Standalone()

	bootstrap_macos_defaults_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_macos_defaults_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_macos_defaults_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if any preference differs from the config (all are still listed)")
	bootstrap_macos_defaultsCmd.AddCommand(bootstrap_macos_defaults_statusCmd)
}
