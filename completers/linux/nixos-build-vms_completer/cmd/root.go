package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixos-build-vms",
	Short: "build NixOS VM test",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nixos-build-vms.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show usage information")
	rootCmd.Flags().Bool("no-out-link", false, "Do not create a symlink to the output path")
	rootCmd.Flags().StringSlice("option", nil, "Set Nix configuration option")
	rootCmd.Flags().Bool("show-trace", false, "Show the stack trace on evaluation errors")

	rootCmd.Flag("option").Nargs = 2

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}
