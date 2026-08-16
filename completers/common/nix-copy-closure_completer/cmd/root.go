package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-copy-closure",
	Short: "copy store objects to or from a remote machine via SSH",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-copy-closure.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("dry-run", false, "Perform a dry run without actually copying")
	rootCmd.Flags().Bool("from", false, "Copy closure from remote machine to local machine")
	rootCmd.Flags().Bool("gzip", false, "Enable compression of the SSH connection")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("include-outputs", false, "Also copy outputs of store derivations included in the closure")
	rootCmd.Flags().Bool("to", false, "Copy closure from local to remote machine")
	rootCmd.Flags().BoolP("use-substitutes", "s", false, "Attempt to download missing store objects on target from substituters")

	rootCmd.MarkFlagsMutuallyExclusive("to", "from")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(), // user@host
	)
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}
