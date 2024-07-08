package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runParallelCmd = &cobra.Command{
	Use:   "run-parallel",
	Short: "run environments in parallel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runParallelCmd).Standalone()

	add_common_flags(runParallelCmd)

	runParallelCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	runParallelCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	runParallelCmd.Flags().StringArray("discover", []string{}, "for Python discovery first try these Python executables (default: [])")
	runParallelCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	runParallelCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	runParallelCmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")
	runParallelCmd.Flags().BoolP("notest", "n", false, "do not run the test commands (default: False)")

	// these two are the same flag; there are two long flags for this
	runParallelCmd.Flags().BoolP("pkg-only", "b", false, "only perform the packaging activity")
	runParallelCmd.Flags().Bool("sdistonly", false, "only perform the packaging activity")
	runParallelCmd.MarkFlagsMutuallyExclusive("pkg-only", "sdistonly")

	runParallelCmd.Flags().String("installpkg", "", "use specified package for installation into venv, instead of packaging the project (default: None)")
	runParallelCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	runParallelCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")
	runParallelCmd.Flags().Bool("skip-pkg-install", false, "skip package installation for this run (default: False)")

	runParallelCmd.Flags().StringP("parallel", "p", "auto", "run tox environments in parallel, the argument controls limit: all, auto - cpu count, some positive number, zero is turn off (default: auto)")
	runParallelCmd.Flags().BoolP("parallel-live", "o", false, "connect to stdout while running environments")
	runParallelCmd.Flags().Bool("parallel-no-spinner", false, "run tox environments in parallel, but don't show the spinner, implies --parallel")


	runParallelCmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	runParallelCmd.Flags().StringArrayS("m", "m", []string{}, "labels to evaluate (default: [])")
	runParallelCmd.Flags().StringArrayS("f", "f", []string{}, "factors to evaluate (passing multiple factors means 'AND', passing this option multiple times means 'OR') (default: [])")
	runParallelCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")

	carapace.Gen(runParallelCmd).FlagCompletion(carapace.ActionMap{
		"parallel":  carapace.ActionValues("all", "auto", "0"),  // or any positive integer
	})

	rootCmd.AddCommand(runParallelCmd)
}
