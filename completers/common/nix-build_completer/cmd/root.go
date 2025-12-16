package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-build",
	Short: "build a Nix expression",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-build.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSlice("arg", nil, "Pass [name] [expr] to Nix functions.")
	rootCmd.Flags().StringSlice("argstr", nil, "Pass [name] [string] to Nix functions")
	rootCmd.Flags().StringP("attr", "A", "", "Attribute to build")
	rootCmd.Flags().Bool("dry-run", false, "Show what store paths would be built or downloaded")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().StringP("include", "I", "", "Include paths")
	rootCmd.Flags().Bool("no-out-link", false, "Do not create a symlink to the output path")
	rootCmd.Flags().StringP("out-link", "o", "", "Change the name of the symlink to the output path")

	rootCmd.Flag("arg").Nargs = 2
	rootCmd.Flag("argstr").Nargs = 2

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"attr": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			arg, _ := rootCmd.Flags().GetStringSlice("arg")
			argstr, _ := rootCmd.Flags().GetStringSlice("argstr")
			opts := nix.AttributeOpts{
				Source:  "default.nix",
				Include: rootCmd.Flag("include").Value.String(),
				Arg:     arg,
				ArgStr:  argstr,
			}
			if len(c.Args) > 0 {
				opts.Source = c.Args[0]
			}
			return nix.ActionAttributes(opts)
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)

}
