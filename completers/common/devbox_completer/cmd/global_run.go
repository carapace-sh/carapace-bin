package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var global_runCmd = &cobra.Command{
	Use:   "run [<script> | <cmd>]",
	Short: "Run a script or command in a shell with access to your packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_runCmd).Standalone()

	global_runCmd.Flags().Bool("all-projects", false, "run command in all projects in the working directory, recursively. If command is not found in any project, it will be skipped.")
	global_runCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_runCmd.PersistentFlags().StringP("env", "e", "", "environment variables to set in the devbox environment")
	global_runCmd.PersistentFlags().String("env-file", "", "path to a file containing environment variables to set in the devbox environment")
	global_runCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_runCmd.Flags().BoolP("list", "l", false, "list all scripts defined in devbox.json")
	global_runCmd.Flags().Bool("omit-nix-env", false, "shell environment will omit the env-vars from print-dev-env")
	global_runCmd.Flags().Bool("pure", false, "if this flag is specified, devbox runs the script in an isolated environment inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained.")
	global_runCmd.Flags().Bool("recompute", false, "recompute environment if needed")
	global_runCmd.Flag("config").Hidden = true
	global_runCmd.Flag("omit-nix-env").Hidden = true
	globalCmd.AddCommand(global_runCmd)

	// TODO environem
	carapace.Gen(global_runCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"env":      env.ActionNameValues(false),
		"env-file": carapace.ActionFiles(),
	})

	// TODO positional completion
}
