package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "halt",
	Short: "halt the machine",
	Long:  "https://www.freedesktop.org/software/systemd/man/latest/halt.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "Force immediate halt/power-off/reboot")
	rootCmd.Flags().Bool("halt", false, "Halt the machine")
	rootCmd.Flags().Bool("help", false, "Show this help")
	rootCmd.Flags().BoolP("no-sync", "n", false, "Don't sync hard disks/storage media before power-off, reboot, or halt")
	rootCmd.Flags().Bool("no-wall", false, "Don't send wall message before halt/power-off/reboot")
	rootCmd.Flags().BoolP("no-wtmp", "d", false, "Don't write wtmp record")
	rootCmd.Flags().BoolP("poweroff", "p", false, "Switch off the machine")
	rootCmd.Flags().Bool("reboot", false, "Reboot the machine")
	rootCmd.Flags().BoolP("wtmp-only", "w", false, "Don't halt/power-off/reboot, just write wtmp record")
}
