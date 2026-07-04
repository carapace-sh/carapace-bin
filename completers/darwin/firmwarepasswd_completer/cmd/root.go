package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "firmwarepasswd",
	Short: "Set firmware password",
	Long:  "https://keith.github.io/xcode-manpages/firmwarepasswd.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("check", false, "Check firmware password")
	rootCmd.Flags().Bool("delete", false, "Delete firmware password")
	rootCmd.Flags().Bool("disable-reset-capability", false, "Disable reset capability")
	rootCmd.Flags().Bool("enable-reset-capability", false, "Enable reset capability")
	rootCmd.Flags().BoolS("h", "h", false, "Display help")
	rootCmd.Flags().Bool("mode", false, "Show mode")
	rootCmd.Flags().Bool("setmode", false, "Set mode")
	rootCmd.Flags().Bool("setpasswd", false, "Set password")
	rootCmd.Flags().Bool("unlockseed", false, "Generate unlock seed")
	rootCmd.Flags().Bool("verify", false, "Verify password")
}
