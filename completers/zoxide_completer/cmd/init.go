package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [OPTIONS] <SHELL",
	Short: "Generate shell configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("cmd", "", "Changes the prefix of the `z` and `zi` commands")
	initCmd.Flags().BoolP("help", "h", false, "Print help")
	initCmd.Flags().String("hook", "", "Changes how often zoxide increments a directory's score")
	initCmd.Flags().Bool("no-cmd", false, "Prevents zoxide from defining the `z` and `zi` commands")
	initCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"hook": carapace.ActionValuesDescribed(
			"none", "Never",
			"prompt", "At every shell prompt",
			"pwd", "Whenever the directory is changed",
		),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionValues("ash", "elvish", "fish", "nushell", "posix", "powershell", "xonsh", "zsh"),
	)
}
