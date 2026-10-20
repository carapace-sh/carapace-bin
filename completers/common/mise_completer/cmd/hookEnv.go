package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var hookEnvCmd = &cobra.Command{
	Use:    "hook-env",
	Short:  "[internal] Update the environment at each prompt",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hookEnvCmd).Standalone()

	hookEnvCmd.Flags().BoolP("force", "f", false, "Skip early exit check")
	hookEnvCmd.Flags().BoolP("help", "h", false, "Print help")
	hookEnvCmd.Flags().BoolP("quiet", "q", false, "Hide warnings such as when a tool is not installed")
	hookEnvCmd.Flags().String("reason", "", "Reason for calling hook-env (e.g., \"precmd\", \"chpwd\")")
	hookEnvCmd.Flags().StringP("shell", "s", "", "Shell type to generate script for")
	hookEnvCmd.Flags().String("shell-pid", "", "")
	hookEnvCmd.Flags().Bool("status", false, "Show \"mise: <TOOL>@<VERSION>\" message when changing directories")
	hookEnvCmd.Flag("reason").Hidden = true
	hookEnvCmd.Flag("shell-pid").Hidden = true
	hookEnvCmd.Flag("status").Hidden = true
	rootCmd.AddCommand(hookEnvCmd)

	carapace.Gen(hookEnvCmd).FlagCompletion(carapace.ActionMap{
		"reason": carapace.ActionValues("precmd", "chpwd"),
		"shell":  action.Shells(),
	})
}
