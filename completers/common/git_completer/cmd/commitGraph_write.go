package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var commitGraph_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write a commit-graph file based on the commits found in packfiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitGraph_writeCmd).Standalone()

	commitGraph_writeCmd.Flags().Bool("append", false, "Include all commits that are present in the existing commit-graph file")
	commitGraph_writeCmd.Flags().Bool("changed-paths", false, "Compute and write information about the paths changed between a commit and its first parent")
	commitGraph_writeCmd.Flags().String("max-new-filters", "", "Generate at most n new Bloom filters")
	commitGraph_writeCmd.Flags().Bool("no-max-new-filters", false, "Do not limit Bloom filters") // TODO verify
	commitGraph_writeCmd.Flags().Bool("no-progress", false, "Turn progress off explicitly")
	commitGraph_writeCmd.Flags().String("object-dir", "", "Use given directory for the location of packfiles and commit-graph file")
	commitGraph_writeCmd.Flags().Bool("progress", false, "Turn progress on explicitly")
	commitGraph_writeCmd.Flags().Bool("reachable", false, "Generate the new commit graph by walking commits starting at all refs.")
	commitGraph_writeCmd.Flags().String("split", "", "Write the commit-graph as a chain of multiple commit-graph files stored in <dir>/info/commit-graphs")
	commitGraph_writeCmd.Flags().Bool("stdin-commits", false, "Generate the new commit graph by walking commits starting at the commits specified in stdin as a list of OIDs in hex, one OID per line")
	commitGraph_writeCmd.Flags().Bool("stdin-packs", false, "Generate the new commit graph by walking objects only in the specified  pack-indexe")
	commitGraphCmd.AddCommand(commitGraph_writeCmd)

	commitGraph_writeCmd.MarkFlagsMutuallyExclusive("reachable", "stdin-packs", "stdin-commits")

	commitGraph_writeCmd.Flag("split").NoOptDefVal = " "

	carapace.Gen(commitGraph_writeCmd).FlagCompletion(carapace.ActionMap{
		"object-dir": carapace.ActionDirectories(),
		"split":      carapace.ActionValues("no-merge", "replace"),
	})
}
