package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packObjectsCmd = &cobra.Command{
	Use:     "pack-objects",
	Short:   "Create a packed archive of objects",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(packObjectsCmd).Standalone()

	packObjectsCmd.Flags().Bool("all", false, "pretend as if all refs under refs/ are specified to be included")
	packObjectsCmd.Flags().Bool("all-progress", false, "progress report is displayed during the object count and compression phases")
	packObjectsCmd.Flags().Bool("all-progress-implied", false, "this is used to imply --all-progress whenever progress display is activated")
	packObjectsCmd.Flags().String("compression", "", "specifies compression level for newly-compressed data in the generated pack")
	packObjectsCmd.Flags().Bool("cruft", false, "packs unreachable objects into a separate \"cruft\" pack")
	packObjectsCmd.Flags().String("cruft-expiration", "", "eliminate objects from the cruft pack older than <approxidate>")
	packObjectsCmd.Flags().Bool("delta-base-offset", false, "express the base object of a delta as offset in the stream")
	packObjectsCmd.Flags().Bool("delta-islands", false, "restrict delta matches based on \"islands\"")
	packObjectsCmd.Flags().String("depth", "", "affect how the objects contained in the pack are stored using delta compression")
	packObjectsCmd.Flags().Bool("exclude-promisor-objects", false, "omit objects that are known to be in the promisor remote")
	packObjectsCmd.Flags().String("filter", "", "omit certain objects from the resulting packfile")
	packObjectsCmd.Flags().Bool("honor-pack-keep", false, "ignore object already in a local pack that has a .keep file,")
	packObjectsCmd.Flags().Bool("include-tag", false, "include unasked-for annotated tags if the object they reference was included in the resulting packfile")
	packObjectsCmd.Flags().Bool("incremental", false, "ignore object already in a pack")
	packObjectsCmd.Flags().String("index-version", "", "force the version for the generated pack index")
	packObjectsCmd.Flags().String("keep-pack", "", "ignore object already in given pack")
	packObjectsCmd.Flags().Bool("keep-unreachable", false, "objects unreachable from the refs in named packs are added to the resulting pack")
	packObjectsCmd.Flags().String("local", "", "object that is borrowed from an alternate object store ")
	packObjectsCmd.Flags().String("max-pack-size", "", "split the output packfile into multiple independent packfiles")
	packObjectsCmd.Flags().String("missing", "", "specify how missing objects are handled")
	packObjectsCmd.Flags().Bool("no-filter", false, "turn off any previous --filter= argument")
	packObjectsCmd.Flags().Bool("no-reuse-delta", false, "do not reuse existing deltas")
	packObjectsCmd.Flags().Bool("no-reuse-object", false, "do not reuse existing object data at all")
	packObjectsCmd.Flags().Bool("no-sparse", false, "disable \"sparse\" algorithm")
	packObjectsCmd.Flags().Bool("non-empty", false, "only create a packed archive if it would contain at least one object")
	packObjectsCmd.Flags().Bool("pack-loose-unreachable", false, "pack unreachable loose objects")
	packObjectsCmd.Flags().Bool("progress", false, "progress status is reported on the standard error stream")
	packObjectsCmd.Flags().BoolS("q", "q", false, "do not to report progress on the standard error stream")
	packObjectsCmd.Flags().Bool("revs", false, "read the revision arguments from the standard input")
	packObjectsCmd.Flags().Bool("shallow", false, "optimize a pack that will be provided to a client with a shallow repository")
	packObjectsCmd.Flags().Bool("sparse", false, "enable \"sparse\" algorithm")
	packObjectsCmd.Flags().Bool("stdin-packs", false, "read the basenames of packfiles from the standard input")
	packObjectsCmd.Flags().Bool("stdout", false, "write the pack contents out to the standard output")
	packObjectsCmd.Flags().Bool("thin", false, "create a \"thin\" pack by omitting the common objects between a sender and a receiver")
	packObjectsCmd.Flags().String("threads", "", "number of threads to spawn when searching for best delta matches")
	packObjectsCmd.Flags().Bool("unpack-unreachable", false, "keep unreachable objects in loose form")
	packObjectsCmd.Flags().Bool("unpacked", false, "limit the objects packed to those that are not already packed")
	packObjectsCmd.Flags().String("window", "", "affect how the objects contained in the pack are stored using delta compression")
	packObjectsCmd.Flags().String("window-memory", "", "additional limit on top of --window")
	rootCmd.AddCommand(packObjectsCmd)

	carapace.Gen(packObjectsCmd).FlagCompletion(carapace.ActionMap{
		"missing": carapace.ActionValuesDescribed(
			"error", "stop with an error if a missing object is encountered",
			"allow-any", "allow object traversal to continue if a missing object is encountered",
			"allow-promisor", "only allow object traversal to continue for EXPECTED promisor missing objects",
		),
	})
}
