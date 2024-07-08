package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	add_common_flags(runCmd)

	runCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	runCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	runCmd.Flags().StringArray("discover", []string{}, "for Python discovery first try these Python executables (default: [])")
	runCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	runCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	runCmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")
	runCmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")

	// these two are the same flag; there are two long flags for this
	runCmd.Flags().BoolP("pkg-only", "b", false, "only perform the packaging activity")
	runCmd.Flags().Bool("sdistonly", false, "only perform the packaging activity")
	runCmd.MarkFlagsMutuallyExclusive("pkg-only", "sdistonly")

	runCmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	runCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	runCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	runCmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")

	runCmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	runCmd.Flags().StringArrayS("m", "m", []string{}, "labels to evaluate (default: [])")
	runCmd.Flags().StringArrayS("f", "f", []string{}, "factors to evaluate (passing multiple factors means 'AND', passing this option multiple times means 'OR') (default: [])")
	runCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")

	rootCmd.AddCommand(runCmd)
}
