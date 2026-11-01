package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteGenerationCmd = &cobra.Command{
	Use:   "delete-generation",
	Short: "delete a previous configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteGenerationCmd).Standalone()

	deleteGenerationCmd.Flags().Bool("allow-reboot", false, "Allow NixOps to reboot the instance if necessary")
	deleteGenerationCmd.Flags().Bool("check", false, "Verify current deployment state")
	deleteGenerationCmd.Flags().StringSlice("exclude", nil, "Operate on all machines except the specified ones")
	deleteGenerationCmd.Flags().Bool("force-reboot", false, "Reboot to activate the new configuration")
	deleteGenerationCmd.Flags().StringSlice("include", nil, "Only operate on the specified machines")
	deleteGenerationCmd.Flags().String("max-concurrent-activate", "", "Max concurrent machine activations")
	deleteGenerationCmd.Flags().String("max-concurrent-copy", "", "Max concurrent nix-copy-closure processes")
	deleteGenerationCmd.Flags().Bool("no-sync", false, "Do not flush buffers to disk")

	rootCmd.AddCommand(deleteGenerationCmd)

	carapace.Gen(deleteGenerationCmd).PositionalCompletion(
		carapace.ActionValues(), // generation
	)
}
