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

	rootCmd.Flags().BoolP("add", "A", false, "Add paths to the Nix store")
	rootCmd.Flags().Bool("add-fixed", false, "Add paths with fixed hashes")
	rootCmd.Flags().String("add-root", "", "Create a symlink to the resulting store path")
	rootCmd.Flags().Bool("check-validity", false, "Check whether paths are valid")
	rootCmd.Flags().Bool("delete", false, "Delete store paths")
	rootCmd.Flags().Bool("dump", false, "Dump a path to stdout")
	rootCmd.Flags().Bool("dump-db", false, "Dump the Nix database")
	rootCmd.Flags().Bool("export", false, "Export store paths to stdout")
	rootCmd.Flags().Bool("gc", false, "Perform garbage collection on the Nix store")
	rootCmd.Flags().Bool("generate-binary-cache-key", false, "Generate a key for a binary cache")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("import", false, "Import store paths from stdout")
	rootCmd.Flags().Bool("init", false, "Initialise the Nix databases")
	rootCmd.Flags().Bool("load-db", false, "Load the Nix database")
	rootCmd.Flags().Bool("optimise", false, "Optimise the Nix store")
	rootCmd.Flags().Bool("print-env", false, "Print the build environment of a derivation")
	rootCmd.Flags().Bool("print-fixed-path", false, "Print the fixed-output store path")
	rootCmd.Flags().BoolP("query", "q", false, "Display information about store paths")
	rootCmd.Flags().BoolP("read-log", "l", false, "Print the build log of a store path")
	rootCmd.Flags().BoolP("realise", "r", false, "Build or fetch store paths")
	rootCmd.Flags().Bool("register-validity", false, "Register validity of paths")
	rootCmd.Flags().Bool("repair-path", false, "Repair a store path")
	rootCmd.Flags().Bool("restore", false, "Restore a path from a NAR archive")
	rootCmd.Flags().Bool("serve", false, "Serve the Nix store over a Unix domain socket")
	rootCmd.Flags().Bool("verify", false, "Verify the Nix store")
	rootCmd.Flags().Bool("verify-path", false, "Verify specific store paths")

	rootCmd.Flags().StringP("binding", "b", "", "Print the value of a store derivation binding")
	rootCmd.Flags().Bool("check", false, "Check whether a derivation is deterministic")
	rootCmd.Flags().Bool("check-contents", false, "Check contents of every valid store path")
	rootCmd.Flags().BoolP("deriver", "d", false, "Print the deriver used to build the paths")
	rootCmd.Flags().Bool("dry-run", false, "Show what store paths would be built or downloaded")
	rootCmd.Flags().BoolP("force-realise", "f", false, "Realise each argument first before querying")
	rootCmd.Flags().Bool("graph", false, "Print references graph in Graphviz dot format")
	rootCmd.Flags().Bool("graphml", false, "Print references graph in GraphML format")
	rootCmd.Flags().Bool("hash", false, "Print SHA-256 hash of path contents")
	rootCmd.Flags().Bool("hash-given", false, "Expect NAR hash and size in the input")
	rootCmd.Flags().Bool("ignore-liveness", false, "Delete paths even if they have remaining referrers")
	rootCmd.Flags().Bool("ignore-unknown", false, "Silently ignore non-derivation paths without substitutes")
	rootCmd.Flags().Bool("include-outputs", false, "Include outputs in the closure computation")
	rootCmd.Flags().String("max-freed", "", "Keep deleting until at least bytes have been freed")
	rootCmd.Flags().Bool("no-output", false, "Suppress printing of output paths")
	rootCmd.Flags().Bool("outputs", false, "Print output paths of store derivations")
	rootCmd.Flags().Bool("print-dead", false, "Print the set of dead store paths")
	rootCmd.Flags().Bool("print-invalid", false, "Print invalid paths instead of throwing an error")
	rootCmd.Flags().Bool("print-live", false, "Print the set of live store paths")
	rootCmd.Flags().Bool("print-roots", false, "Print the set of GC roots")
	rootCmd.Flags().Bool("recursive", false, "Use recursive hashing mode for directories")
	rootCmd.Flags().Bool("references", false, "Print immediate dependencies")
	rootCmd.Flags().Bool("referrers", false, "Print store paths that refer to the given paths")
	rootCmd.Flags().Bool("referrers-closure", false, "Print closure under the referrers relation")
	rootCmd.Flags().Bool("repair", false, "Repair corrupted or missing store paths")
	rootCmd.Flags().BoolP("requisites", "R", false, "Print closure of store paths")
	rootCmd.Flags().Bool("reregister", false, "Register even if already marked as valid")
	rootCmd.Flags().Bool("resolve", false, "Resolve symlinks to store paths")
	rootCmd.Flags().Bool("roots", false, "Print GC roots pointing to the given paths")
	rootCmd.Flags().Bool("size", false, "Print size in bytes of path contents")
	rootCmd.Flags().Bool("tree", false, "Print references graph as ASCII tree")
	rootCmd.Flags().BoolP("use-output", "u", false, "Apply the query to the output path of the derivation")
	rootCmd.Flags().Bool("valid-derivers", false, "Print derivation files that produce said paths")
	rootCmd.Flags().Bool("write", false, "Allow write operations via the serve protocol")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add-root": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}
