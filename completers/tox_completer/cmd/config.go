package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "show tox configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	add_common_flags(configCmd)

	configCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	configCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	configCmd.Flags().StringArray("discover", []string{}, "for Python discovery first try these Python executables (default: [])")
	configCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	configCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	configCmd.Flags().StringS("k", "k", "", "list just configuration keys specified (default: [])")
	configCmd.Flags().Bool("core", false, "show core options (by default is hidden unless -e ALL is passed) (default: False)")
	configCmd.Flags().StringP("skip-missing-interpreters", "s", "config", "don't fail tests for missing interpreters: {config,true,false} choice (default: config)")
	configCmd.Flags().Bool("develop", false, "install package in development mode (default: False)")
	configCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")

	configCmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	configCmd.Flags().StringArrayS("m", "m", []string{}, "labels to evaluate (default: [])")
	configCmd.Flags().StringArrayS("f", "f", []string{}, "factors to evaluate (passing multiple factors means 'AND', passing this option multiple times means 'OR') (default: [])")
	configCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"skip-missing-interpreters":  carapace.ActionValues("config", "true", "false"),
		"e":  tox.ActionEnvironments().UniqueList(","),
	})

	rootCmd.AddCommand(configCmd)
}
