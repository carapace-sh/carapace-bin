package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-shell",
	Short: "start an interactive shell based on a Nix expression",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-shell.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringSlice("arg", nil, "Pass [name] [expr] to Nix functions.")
	rootCmd.Flags().StringSlice("argstr", nil, "Pass [name] [string] to Nix functions")
	rootCmd.Flags().String("command", "", "Run the command in an interactive shell")
	rootCmd.Flags().String("exclude", "", "Do not build any dependencies whose store path matches the regexp")
	rootCmd.Flags().StringP("expr", "E", "", "Nix expression to build")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().StringP("include", "I", "", "Include paths")
	rootCmd.Flags().String("keep", "", "Environment variable to keep in --pure")
	rootCmd.Flags().StringSliceP("packages", "p", nil, "Packages to add from nixpkgs")
	rootCmd.Flags().Bool("pure", false, "Clear most of the environment for the shell")
	rootCmd.Flags().String("run", "", "Run the command in a non-interactive shell")

	rootCmd.Flag("arg").Nargs = 2
	rootCmd.Flag("argstr").Nargs = 2

	rootCmd.MarkFlagsMutuallyExclusive("packages", "expr")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"packages": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			opts := nix.AttributeOpts{Source: "<nixpkgs>", Include: rootCmd.Flag("include").Value.String()}
			return nix.ActionAttributes(opts)
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}
