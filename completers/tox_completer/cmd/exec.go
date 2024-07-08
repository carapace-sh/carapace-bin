package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Aliases: []string{"e"},
	Short: "execute an arbitrary command within a tox environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	add_common_flags(execCmd)

	execCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	execCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 2416698050)")
	execCmd.Flags().StringArray("discover", []string{}, "for Python discovery first try these Python executables (default: [])")
	execCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	execCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	execCmd.Flags().StringS("e", "e", "", "environment to run (default: py)")
	execCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")
	execCmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")
	execCmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")

	// these two are the same flag; there are two long flags for this
	execCmd.Flags().BoolP("pkg-only", "b", false, "only perform the packaging activity")
	execCmd.Flags().Bool("sdistonly", false, "only perform the packaging activity")
	execCmd.MarkFlagsMutuallyExclusive("pkg-only", "sdistonly")

	execCmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	execCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	execCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	execCmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"e":  tox.ActionEnvironments(),
	})

	// The exec subcommand is designed to run an arbitrary command after '--',
	// so here we can bridge to a new base completer.
	carapace.Gen(execCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)

	rootCmd.AddCommand(execCmd)
}
