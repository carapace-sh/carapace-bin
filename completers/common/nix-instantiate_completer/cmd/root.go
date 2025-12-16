package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-instantiate",
	Short: "instantiate store derivations from Nix expression",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-instantiate.html",
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

	rootCmd.Flags().String("add-root", "", "Create a symlink to the output path")
	rootCmd.Flags().StringSlice("arg", nil, "Pass [name] [expr] to Nix functions.")
	rootCmd.Flags().StringSlice("argstr", nil, "Pass [name] [string] to Nix functions")
	rootCmd.Flags().StringP("attr", "A", "", "Attribute to build")
	rootCmd.Flags().Bool("dry-run", false, "Show what store paths would be built or downloaded")
	rootCmd.Flags().Bool("eval", false, "Evaluate the input files, and print the resulting values on standard output")
	rootCmd.Flags().StringP("expr", "E", "", "The expression to evaluate")
	rootCmd.Flags().Bool("find-file", false, "Lookup files in the nix search path")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().StringP("include", "I", "", "Include paths")
	rootCmd.Flags().Bool("json", false, "Print the result of --eval as JSON")
	rootCmd.Flags().Bool("parse", false, "Just parse the input files, and print their abstract syntax trees on standard output in ATerm format.")
	rootCmd.Flags().Bool("read-write-mode", false, "Perform evaluation in read/write mode")
	rootCmd.Flags().Bool("strict", false, "Recursively evaluate list elements and attributes.")
	rootCmd.Flags().Bool("xml", false, "Print the result of --eval as XML")

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
