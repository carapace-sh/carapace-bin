package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a disk image to a target",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()
	restoreCmd.Flags().Bool("erase", false, "Erase the target before restoring")
	restoreCmd.Flags().String("format", "", "Format for the target volume")
	restoreCmd.Flags().Bool("noverify", false, "Skip verification of the source image")
	restoreCmd.Flags().Bool("puppetstrings", false, "Print puppet strings for scriptability")
	restoreCmd.Flags().String("source", "", "Source image to restore from")
	restoreCmd.Flags().String("target", "", "Target device to restore to")
	rootCmd.AddCommand(restoreCmd)
	carapace.Gen(restoreCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("HFS+", "APFS"),
		"source": carapace.ActionFiles(),
		"target": carapace.ActionFiles(),
	})
}
