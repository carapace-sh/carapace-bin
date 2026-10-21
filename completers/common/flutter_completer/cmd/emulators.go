package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var emulatorsCmd = &cobra.Command{
	Use:     "emulators",
	Aliases: []string{"emulator"},
	Short:   "List, launch and create emulators",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emulatorsCmd).Standalone()

	emulatorsCmd.Flags().Bool("cold", false, "Used with the \"--launch\" flag to cold boot the emulator instance (Android only).")
	emulatorsCmd.Flags().Bool("create", false, "Creates a new Android emulator based on a Pixel device.")
	emulatorsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	emulatorsCmd.Flags().String("launch", "", "The full or partial ID of the emulator to launch.")
	emulatorsCmd.Flags().String("name", "", "Used with the \"--create\" flag. Specifies a name for the emulator being created.")
	rootCmd.AddCommand(emulatorsCmd)

	carapace.Gen(emulatorsCmd).FlagCompletion(carapace.ActionMap{
		"launch": action.ActionEmulators(),
	})
}
