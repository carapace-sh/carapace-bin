package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-store",
	Short: "manipulate or query the Nix store",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-store.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("add", false, "Add paths to the Nix store")
	rootCmd.Flags().Bool("add-fixed", false, "Add paths with fixed hashes")
	rootCmd.Flags().String("add-root", "", "Create a symlink to the resulting store path")
	rootCmd.Flags().Bool("clear-failed-paths", false, "Clear all failed paths")
	rootCmd.Flags().Bool("delete", false, "Delete store paths")
	rootCmd.Flags().Bool("dump", false, "Dump a path to stdout")
	rootCmd.Flags().Bool("dump-db", false, "Dump the Nix database")
	rootCmd.Flags().Bool("export", false, "Export store paths to stdout")
	rootCmd.Flags().Bool("gc", false, "Perform garbage collection on the Nix store")
	rootCmd.Flags().Bool("generate-binary-cache-key", false, "Generate a key for a binary cache")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("import", false, "Import store paths from stdout")
	rootCmd.Flags().Bool("load-db", false, "Load the Nix database")
	rootCmd.Flags().Bool("optimise", false, "Optimise the Nix store")
	rootCmd.Flags().Bool("print-env", false, "Print the build environment of a derivation")
	rootCmd.Flags().BoolP("query", "q", false, "Display information about store paths")
	rootCmd.Flags().Bool("query-failed-paths", false, "Query all failed paths")
	rootCmd.Flags().BoolP("read-log", "l", false, "Print the build log of a store path")
	rootCmd.Flags().BoolP("realise", "r", false, "Build or fetch store paths")
	rootCmd.Flags().Bool("repair-path", false, "Repair a store path")
	rootCmd.Flags().Bool("restore", false, "Restore a path from a NAR archive")
	rootCmd.Flags().Bool("serve", false, "Serve the Nix store over a Unix domain socket")
	rootCmd.Flags().Bool("verify", false, "Verify the Nix store")
	rootCmd.Flags().Bool("verify-path", false, "Verify specific store paths")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-root": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}