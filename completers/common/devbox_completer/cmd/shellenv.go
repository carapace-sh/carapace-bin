package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var shellenvCmd = &cobra.Command{
	Use:   "shellenv",
	Short: "Print shell commands that create a Devbox Environment in the shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellenvCmd).Standalone()

	shellenvCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	shellenvCmd.PersistentFlags().StringP("env", "e", "", "environment variables to set in the devbox environment")
	shellenvCmd.PersistentFlags().String("env-file", "", "path to a file containing environment variables to set in the devbox environment")
	shellenvCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	shellenvCmd.Flags().Bool("init-hook", false, "runs init hook after exporting shell environment")
	shellenvCmd.Flags().Bool("install", false, "install packages before exporting shell environment")
	shellenvCmd.Flags().Bool("no-refresh-alias", false, "by default, devbox will add refresh alias to the environmentUse this flag to disable this behavior.")
	shellenvCmd.Flags().Bool("omit-nix-env", false, "shell environment will omit the env-vars from print-dev-env")
	shellenvCmd.Flags().Bool("preserve-path-stack", false, "preserves existing PATH order if this project's environment is already in PATH. Useful if you want to avoid overshadowing another devbox project that is already active")
	shellenvCmd.Flags().Bool("pure", false, "if this flag is specified, devbox creates an isolated environment inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained.")
	shellenvCmd.Flags().BoolP("recompute", "r", false, "Recompute environment if needed")
	shellenvCmd.Flag("no-refresh-alias").Hidden = true
	shellenvCmd.Flag("omit-nix-env").Hidden = true
	shellenvCmd.Flag("preserve-path-stack").Hidden = true
	rootCmd.AddCommand(shellenvCmd)

	// TODO environment
	carapace.Gen(shellenvCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"env":      env.ActionNameValues(false),
		"env-file": carapace.ActionFiles(),
	})
}
