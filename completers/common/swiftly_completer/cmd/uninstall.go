package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall <toolchain>...",
	Short: "Remove installed toolchains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().Bool("assume-yes", false, "Disable confirmation prompts by assuming 'yes'")
	uninstallCmd.Flags().String("format", "text", "Output format (text|json)")
	uninstallCmd.Flags().BoolP("help", "h", false, "Show help information")

	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"format": actionFormat(),
	})

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		actionToolchains(),
	)
}
