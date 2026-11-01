package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

// AddEvaluationFlags adds common evaluation flags
// (MixEvalArgs in nix source).
func AddEvaluationFlags(cmd *cobra.Command) {
	cmd.Flags().StringSlice("arg", nil, "Pass the value expr as the argument name to Nix functions")
	cmd.Flags().StringSlice("argstr", nil, "Pass the string string as the argument name to Nix functions")
	cmd.Flags().Bool("debugger", false, "Start an interactive environment if evaluation fail")
	cmd.Flags().String("eval-store", "", "The Nix store to use for evaluations")
	cmd.Flags().Bool("impure", false, "Allow access to mutable paths and repositories")
	cmd.Flags().BoolP("include", "I", false, "Add path to the list of locations used to look up <...> file names")
	cmd.Flags().String("override-flake", "", "Override the flake registries, redirecting original-ref to resolved-ref")

	cmd.Flag("arg").Nargs = 2
	cmd.Flag("argstr").Nargs = 2

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"include":        carapace.ActionFiles(),
		"override-flake": nix.ActionFlakeRefs(),
	})
}
