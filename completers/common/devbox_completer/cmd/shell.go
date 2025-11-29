package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a new shell with access to your packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()

	shellCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	shellCmd.PersistentFlags().StringP("env", "e", "", "environment variables to set in the devbox environment")
	shellCmd.PersistentFlags().String("env-file", "", "path to a file containing environment variables to set in the devbox environment")
	shellCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	shellCmd.Flags().Bool("omit-nix-env", false, "shell environment will omit the env-vars from print-dev-env")
	shellCmd.Flags().Bool("print-env", false, "print script to setup shell environment")
	shellCmd.Flags().Bool("pure", false, "if this flag is specified, devbox creates an isolated shell inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained.")
	shellCmd.Flags().Bool("recompute", false, "recompute environment if needed")
	shellCmd.Flag("omit-nix-env").Hidden = true
	rootCmd.AddCommand(shellCmd)

	// TODO environment
	carapace.Gen(shellCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"env":      env.ActionNameValues(false),
		"env-file": carapace.ActionFiles(),
	})

	carapace.Gen(shellCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	// TODO dash completion - if possible?
}
