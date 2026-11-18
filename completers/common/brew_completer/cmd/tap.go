package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
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
	tapCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	tapCmd.Flags().Bool("eval-all", false, "Evaluate all the formulae, casks and aliases in the new tap to check validity. Implied if `HOMEBREW_EVAL_ALL` is set.")
	tapCmd.Flags().BoolP("force", "f", false, "Force install core taps even under API mode.")
	tapCmd.Flags().Bool("force-auto-update", false, "Auto-update tap even if it is not hosted on GitHub. By default, only taps hosted on GitHub are auto-updated (for performance reasons).")
	tapCmd.Flags().BoolP("help", "h", false, "Show this message.")
	tapCmd.Flags().Bool("no-force-auto-update", false, "Auto-update tap even if it is not hosted on GitHub. By default, only taps hosted on GitHub are auto-updated (for performance reasons).")
	tapCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	tapCmd.Flags().Bool("repair", false, "Migrate tapped formulae from symlink-based to directory-based structure.")
	tapCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(tapCmd)

	carapace.Gen(tapCmd).PositionalCompletion(
		brew.ActionInstalledTaps(),
		carapace.ActionValues(),
	)
}
