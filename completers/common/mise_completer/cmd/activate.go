// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Print the script to activate mise in an interactive shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(activateCmd).Standalone()

	activateCmd.Flags().BoolP("help", "h", false, "Print help")
	activateCmd.Flags().Bool("no-hook-env", false, "Do not run hook-env automatically")
	activateCmd.Flags().BoolP("quiet", "q", false, "Suppress non-error messages from the hook that runs at each prompt")
	activateCmd.Flags().StringP("shell", "s", "", "Shell type to generate the script for")
	activateCmd.Flags().Bool("shims", false, "Put the shims directory on PATH instead of updating PATH at each prompt")
	activateCmd.Flags().Bool("status", false, "Show \"mise: <TOOL>@<VERSION>\" message when changing directories")
	activateCmd.Flag("shell").Hidden = true
	activateCmd.Flag("status").Hidden = true
	rootCmd.AddCommand(activateCmd)

	carapace.Gen(activateCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "elvish", "fish", "nu", "xonsh", "zsh", "pwsh", "powershell"),
	})

	carapace.Gen(activateCmd).PositionalCompletion(
		carapace.ActionValues("bash", "elvish", "fish", "nu", "xonsh", "zsh", "pwsh", "powershell"),
	)
}
