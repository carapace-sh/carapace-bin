package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsckCmd = &cobra.Command{
	Use:     "fsck",
	Short:   "Verifies the connectivity and validity of the objects in the database",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(fsckCmd).Standalone()

	fsckCmd.Flags().Bool("cache", false, "consider any object recorded in the index also as a head node for an unreachability trace")
	fsckCmd.Flags().Bool("connectivity-only", false, "check only the connectivity of reachable objects")
	fsckCmd.Flags().Bool("dangling", false, "print objects that exist but that are never directly used")
	fsckCmd.Flags().Bool("full", false, "check not just objects in GIT_OBJECT_DIRECTORY")
	fsckCmd.Flags().Bool("lost-found", false, "write dangling objects into .git/lost-found/commit/ or .git/lost-found/other/")
	fsckCmd.Flags().Bool("name-objects", false, "also display a name that describes how objects are reachable")
	fsckCmd.Flags().Bool("no-dangling", false, "do not print objects that exist but that are never directly used")
	fsckCmd.Flags().Bool("no-full", false, "check only objects in GIT_OBJECT_DIRECTORY")
	fsckCmd.Flags().Bool("no-progress", false, "do no report progress status")
	fsckCmd.Flags().Bool("no-reflogs", false, "do not consider commits that are referenced only by an entry in a reflog to be reachable")
	fsckCmd.Flags().Bool("progress", false, "report progress status")
	fsckCmd.Flags().Bool("root", false, "report root nodes")
	fsckCmd.Flags().Bool("strict", false, "enable more strict checking")
	fsckCmd.Flags().Bool("tags", false, "report tags")
	fsckCmd.Flags().Bool("unreachable", false, "print out objects that exist but that aren’t reachable from any of the reference nodes")
	fsckCmd.Flags().Bool("verbose", false, "be chatty")
	rootCmd.AddCommand(fsckCmd)

	// TODO positional completion
}
