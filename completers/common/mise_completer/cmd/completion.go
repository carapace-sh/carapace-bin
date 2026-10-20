package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:     "completion",
	Short:   "Generate shell completions",
	Aliases: []string{"complete", "completions"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	completionCmd.Flags().Bool("force", false, "Replace a file at the target path that mise did not write")
	completionCmd.Flags().BoolP("help", "h", false, "Print help")
	completionCmd.Flags().Bool("include-bash-completion-lib", false, "No longer has any effect; kept so older setup scripts keep working")
	completionCmd.Flags().Bool("install", false, "Install the script where this shell looks for it, instead of printing it")
	completionCmd.Flags().StringP("shell", "s", "", "Shell type to generate completions for")
	completionCmd.Flags().String("tool", "", "Generate completions for a tool installed with the `packslip:` backend")
	completionCmd.Flags().Bool("usage", false, "No longer has any effect; kept so older setup scripts keep working")
	completionCmd.Flag("include-bash-completion-lib").Hidden = true
	completionCmd.Flag("shell").Hidden = true
	completionCmd.Flag("usage").Hidden = true
	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "fish", "powershell", "zsh", "pwsh"),
	})

	carapace.Gen(completionCmd).PositionalCompletion(
		carapace.ActionValues("bash", "fish", "powershell", "zsh", "pwsh"),
	)
}
