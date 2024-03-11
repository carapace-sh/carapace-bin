package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tapCmd = &cobra.Command{
	Use:     "tap",
	Short:   "Tap a formula repository",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tapCmd).Standalone()

	tapCmd.Flags().Bool("custom-remote", false, "Install or change a tap with a custom remote. Useful for mirrors.")
	tapCmd.Flags().Bool("debug", false, "Display any debugging information.")
	tapCmd.Flags().Bool("eval-all", false, "Evaluate all the formulae, casks and aliases in the new tap to check validity. Implied if `HOMEBREW_EVAL_ALL` is set.")
	tapCmd.Flags().Bool("force", false, "Force install core taps even under API mode.")
	tapCmd.Flags().Bool("force-auto-update", false, "Auto-update tap even if it is not hosted on GitHub. By default, only taps hosted on GitHub are auto-updated (for performance reasons).")
	tapCmd.Flags().Bool("help", false, "Show this message.")
	tapCmd.Flags().Bool("no-force-auto-update", false, "Auto-update tap even if it is not hosted on GitHub. By default, only taps hosted on GitHub are auto-updated (for performance reasons).")
	tapCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	tapCmd.Flags().Bool("repair", false, "Migrate tapped formulae from symlink-based to directory-based structure.")
	tapCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(tapCmd)
}
