package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showConsoleOutputCmd = &cobra.Command{
	Use:   "show-console-output",
	Short: "print the machine console output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showConsoleOutputCmd).Standalone()

	showConsoleOutputCmd.Flags().Bool("allow-reboot", false, "Allow NixOps to reboot the instance if necessary")
	showConsoleOutputCmd.Flags().Bool("check", false, "Verify current deployment state")
	showConsoleOutputCmd.Flags().StringSlice("exclude", nil, "Operate on all machines except the specified ones")
	showConsoleOutputCmd.Flags().Bool("force-reboot", false, "Reboot to activate the new configuration")
	showConsoleOutputCmd.Flags().StringSlice("include", nil, "Only operate on the specified machines")
	showConsoleOutputCmd.Flags().String("max-concurrent-activate", "", "Max concurrent machine activations")
	showConsoleOutputCmd.Flags().String("max-concurrent-copy", "", "Max concurrent nix-copy-closure processes")
	showConsoleOutputCmd.Flags().Bool("no-sync", false, "Do not flush buffers to disk")

	rootCmd.AddCommand(showConsoleOutputCmd)

	carapace.Gen(showConsoleOutputCmd).PositionalCompletion(
		carapace.ActionValues(), // machine
	)
}
