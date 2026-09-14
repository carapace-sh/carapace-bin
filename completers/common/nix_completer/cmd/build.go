package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "build a derivation or fetch a store path",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	buildCmd.Flags().Bool("no-link", false, "Do not create symlinks to the build results.")
	buildCmd.Flags().StringP("out-link", "o", "", "Use path as prefix for the symlinks to the build results.")
	buildCmd.Flags().Bool("print-out-paths", false, "Print the resulting output paths")
	buildCmd.Flags().String("profile", "", "The profile to update.")
	buildCmd.Flags().Bool("rebuild", false, "Rebuild an already built package and compare the result to the existing store paths.")
	buildCmd.Flags().Bool("stdin", false, "Read installables from the standard input")

	common.AddEvaluationFlags(buildCmd)
	common.AddFlakeFlags(buildCmd)
	common.AddInterpretationFlags(buildCmd)
	common.AddLoggingFlags(buildCmd)

	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"out-link": carapace.ActionFiles(),
		"profile":  carapace.ActionFiles(),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if f := buildCmd.Flag("file"); f.Changed {
				arg, _ := rootCmd.Flags().GetStringSlice("arg")
				argstr, _ := rootCmd.Flags().GetStringSlice("argstr")
				opts := nix.AttributeOpts{
					Source:  "default.nix",
					Include: buildCmd.Flag("include").Value.String(),
					Arg:     arg,
					ArgStr:  argstr,
				}
				return nix.ActionAttributes(opts)
			}

			return carapace.Batch(
				carapace.ActionDirectories(),
				nix.ActionFlakeRefs(),
			).ToA()
		}),
	)
}
