package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repository_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a remote repository into the given path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_cloneCmd).Standalone()

	repository_cloneCmd.Flags().Bool("bare", false, "Clone without files, only fetch latest revision tree")
	repository_cloneCmd.Flags().String("branch", "", "Optional branch to sync (shorthand for a full revision specifier)")
	repository_cloneCmd.Flags().String("dependency-depth-limit", "0", "Maximum dependency traversal depth (0 means unlimited)")
	repository_cloneCmd.Flags().Bool("dependency-recursive", false, "Follow transitive dependencies recursively during dependency-based clone")
	repository_cloneCmd.Flags().StringSlice("dependency-tag", nil, "Tags to filter dependencies by during dependency-based clone")
	repository_cloneCmd.Flags().Bool("direct-file-io", false, "Use direct file I/O instead of memory mapping files")
	repository_cloneCmd.Flags().Bool("direct-file-write", false, "Write directly to the destination file instead of write to a temporary file and move into place")
	repository_cloneCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_cloneCmd.Flags().String("layer", "", "Layer to add")
	repository_cloneCmd.Flags().String("layer-metadata", "", "Metadata key to link layer revisions with")
	repository_cloneCmd.Flags().Bool("no-tracking", false, "Clone without local repository tracking (memory-only stores)")
	repository_cloneCmd.Flags().String("prefetch", "", "File containing list of files to prefetch")
	repository_cloneCmd.Flags().String("revision", "", "Optional revision to sync")
	repository_cloneCmd.Flags().StringSlice("root-file", nil, "Root files for dependency-based selective clone (only clone these files and their dependencies)")
	repository_cloneCmd.Flags().String("shared-store-path", "", "Use this path rather than the system default as the shared store location")
	repository_cloneCmd.Flags().Bool("use-shared-store", false, "Use the shared store rather than create a local immutable store")
	repository_cloneCmd.Flags().String("view", "", "Optional client side view filter file")
	repository_cloneCmd.Flags().Bool("virtual", false, "Clone virtually using split-write filesystem")
	repositoryCmd.AddCommand(repository_cloneCmd)

	carapace.Gen(repository_cloneCmd).FlagCompletion(carapace.ActionMap{
		"branch":            action.ActionBranches(repository_cloneCmd),
		"layer":             carapace.ActionValues(),
		"layer-metadata":    carapace.ActionValues(),
		"prefetch":          carapace.ActionFiles(),
		"revision":          action.ActionRevisions(repository_cloneCmd),
		"root-file":         carapace.ActionFiles(),
		"shared-store-path": carapace.ActionDirectories(),
		"view":              carapace.ActionFiles(),
	})

	carapace.Gen(repository_cloneCmd).PositionalCompletion(
		carapace.ActionValues(), // remote url (e.g. lore://host:port/repo)
		carapace.ActionDirectories(),
	)
}
