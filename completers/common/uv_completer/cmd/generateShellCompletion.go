package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateShellCompletionCmd = &cobra.Command{
	Use:     "generate-shell-completion",
	Short:   "Generate shell completion",
	Aliases: []string{"--generate-shell-completion"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateShellCompletionCmd).Standalone()

	generateShellCompletionCmd.Flags().String("cache-dir", "", "")
	generateShellCompletionCmd.Flags().String("color", "", "")
	generateShellCompletionCmd.Flags().String("config-file", "", "")
	generateShellCompletionCmd.Flags().BoolP("help", "h", false, "")
	generateShellCompletionCmd.Flags().Bool("native-tls", false, "")
	generateShellCompletionCmd.Flags().BoolP("no-cache", "n", false, "")
	generateShellCompletionCmd.Flags().Bool("no-config", false, "")
	generateShellCompletionCmd.Flags().Bool("no-progress", false, "")
	generateShellCompletionCmd.Flags().Bool("no-python-downloads", false, "")
	generateShellCompletionCmd.Flags().Bool("offline", false, "")
	generateShellCompletionCmd.Flags().String("python-preference", "", "")
	generateShellCompletionCmd.Flags().CountP("quiet", "q", "")
	generateShellCompletionCmd.Flags().CountP("verbose", "v", "")
	generateShellCompletionCmd.Flags().BoolP("version", "V", false, "")
	generateShellCompletionCmd.Flag("cache-dir").Hidden = true
	generateShellCompletionCmd.Flag("color").Hidden = true
	generateShellCompletionCmd.Flag("config-file").Hidden = true
	generateShellCompletionCmd.Flag("help").Hidden = true
	generateShellCompletionCmd.Flag("native-tls").Hidden = true
	generateShellCompletionCmd.Flag("no-cache").Hidden = true
	generateShellCompletionCmd.Flag("no-config").Hidden = true
	generateShellCompletionCmd.Flag("no-progress").Hidden = true
	generateShellCompletionCmd.Flag("no-python-downloads").Hidden = true
	generateShellCompletionCmd.Flag("offline").Hidden = true
	generateShellCompletionCmd.Flag("python-preference").Hidden = true
	generateShellCompletionCmd.Flag("quiet").Hidden = true
	generateShellCompletionCmd.Flag("verbose").Hidden = true
	generateShellCompletionCmd.Flag("version").Hidden = true
	rootCmd.AddCommand(generateShellCompletionCmd)
	carapace.Gen(generateShellCompletionCmd).PositionalCompletion(carapace.ActionValues(
		"bash",
		"elvish",
		"fish",
		"nushell",
		"powershell",
		"zsh",
	))
}
