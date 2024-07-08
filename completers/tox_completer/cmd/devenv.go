package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/spf13/cobra"
)

var devenvCmd = &cobra.Command{
	Use:   "devenv",
	Short: "sets up a development environment at ENVDIR based on the tox configuration specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devenvCmd).Standalone()

	add_common_flags(devenvCmd)

	devenvCmd.Flags().String("result-json", "", "write a JSON file with detailed information about all commands and results involved (default: None)")
	devenvCmd.Flags().String("hashseed", "", "set PYTHONHASHSEED to SEED before running commands. Defaults to a random integer in the range [1, 4294967295] ([1, 1024] on Windows). Passing 'notset' suppresses this behavior. (default: 264197440)")
	devenvCmd.Flags().StringArray("discover", []string{}, "for Python discovery first try these Python executables (default: [])")
	devenvCmd.Flags().Bool("list-dependencies", false, "list the dependencies installed during environment setup (default: False)")
	devenvCmd.Flags().Bool("no-list-dependencies", false, "never list the dependencies installed during environment setup (default: True)")
	devenvCmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")
	devenvCmd.Flags().String("skip-env", "", "exclude all environments selected that match this regular expression (default: '')")
	devenvCmd.Flags().Bool("no-recreate-pkg", false, "if recreate is set do not recreate packaging tox environment(s) (default: False)")

	carapace.Gen(devenvCmd).FlagCompletion(carapace.ActionMap{
		"e":  tox.ActionEnvironments(),
	})

	// NOTE: this command has a single optional positional argument:
	// `path (default: venv)`
	// This is arbitrary, so no completion.

	rootCmd.AddCommand(devenvCmd)
}
