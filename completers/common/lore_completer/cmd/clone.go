package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a remote repository into the given path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().Bool("bare", false, "Clone without files, only fetch latest revision tree")
	cloneCmd.Flags().String("branch", "", "Optional branch to sync (shorthand for a full revision specifier)")
	cloneCmd.Flags().String("dependency-depth-limit", "0", "Maximum dependency traversal depth (0 means unlimited)")
	cloneCmd.Flags().Bool("dependency-recursive", false, "Follow transitive dependencies recursively during dependency-based clone")
	cloneCmd.Flags().StringSlice("dependency-tag", nil, "Tags to filter dependencies by during dependency-based clone")
	cloneCmd.Flags().Bool("direct-file-io", false, "Use direct file I/O instead of memory mapping files")
	cloneCmd.Flags().Bool("direct-file-write", false, "Write directly to the destination file instead of write to a temporary file and move into place")
	cloneCmd.Flags().BoolP("help", "h", false, "Print help")
	cloneCmd.Flags().String("layer", "", "Layer to add")
	cloneCmd.Flags().String("layer-metadata", "", "Metadata key to link layer revisions with")
	cloneCmd.Flags().Bool("no-tracking", false, "Clone without local repository tracking (memory-only stores)")
	cloneCmd.Flags().String("prefetch", "", "File containing list of files to prefetch")
	cloneCmd.Flags().String("revision", "", "Optional revision to sync")
	cloneCmd.Flags().StringSlice("root-file", nil, "Root files for dependency-based selective clone (only clone these files and their dependencies)")
	cloneCmd.Flags().String("shared-store-path", "", "Use this path rather than the system default as the shared store location")
	cloneCmd.Flags().Bool("use-shared-store", false, "Use the shared store rather than create a local immutable store")
	cloneCmd.Flags().String("view", "", "Optional client side view filter file")
	cloneCmd.Flags().Bool("virtual", false, "Clone virtually using split-write filesystem")
	rootCmd.AddCommand(cloneCmd)
}
