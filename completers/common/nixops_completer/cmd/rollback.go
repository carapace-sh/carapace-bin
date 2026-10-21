package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "roll back to a previous configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollbackCmd).Standalone()

	rollbackCmd.Flags().Bool("allow-reboot", false, "Allow NixOps to reboot the instance if necessary")
	rollbackCmd.Flags().Bool("check", false, "Verify current deployment state")
	rollbackCmd.Flags().StringSlice("exclude", nil, "Operate on all machines except the specified ones")
	rollbackCmd.Flags().Bool("force-reboot", false, "Reboot to activate the new configuration")
	rollbackCmd.Flags().StringSlice("include", nil, "Only operate on the specified machines")
	rollbackCmd.Flags().String("max-concurrent-activate", "", "Max concurrent machine activations")
	rollbackCmd.Flags().String("max-concurrent-copy", "", "Max concurrent nix-copy-closure processes")
	rollbackCmd.Flags().Bool("no-sync", false, "Do not flush buffers to disk")

	rootCmd.AddCommand(rollbackCmd)

	carapace.Gen(rollbackCmd).PositionalCompletion(
		carapace.ActionValues(), // generation
	)
}
