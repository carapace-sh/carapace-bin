package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-build",
	Short: "build a Nix expression",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-build.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("attr", "A", "", "Attribute to build")
	rootCmd.Flags().Bool("dry-run", false, "Show what store paths would be built or downloaded")
	rootCmd.Flags().StringP("include", "I", "", "Include paths")
	rootCmd.Flags().Bool("no-out-link", false, "Do not create a symlink to the output path")
	rootCmd.Flags().StringP("out-link", "o", "", "Change the name of the symlink to the output path")
	// TODO support --arg and --argstr

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"attr": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			opts := nix.AttributeOpts{Source: "default.nix", Include: rootCmd.Flag("include").Value.String()}
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
