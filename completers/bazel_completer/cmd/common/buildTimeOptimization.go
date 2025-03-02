package common

import (
	"github.com/spf13/cobra"
)

func AddBuildTimeOptimizationFlags(cmd *cobra.Command) {
	cmd.Flags().String("gc_thrashing_limits", "", "Limits which, if reached, cause GcThrashingDetector to crash Bazel with an OOM")
	cmd.Flags().Bool("heuristically_drop_nodes", false, "If true, Blaze will remove FileState and DirectoryListingState nodes after related File and DirectoryListing node is done to save memory")
	cmd.Flags().Bool("incompatible_do_not_split_linking_cmdline", false, "When true, Bazel no longer modifies command line flags used for linking, and also doesn't selectively decide which flags go to the param file and which don't")
	cmd.Flags().Bool("keep_state_after_build", false, "If false, Blaze will discard the inmemory state from this build when the build finishes")
	cmd.Flags().Bool("no-heuristically_drop_nodes", false, "If true, Blaze will remove FileState and DirectoryListingState nodes after related File and DirectoryListing node is done to save memory")
	cmd.Flags().Bool("no-incompatible_do_not_split_linking_cmdline", false, "When true, Bazel no longer modifies command line flags used for linking, and also doesn't selectively decide which flags go to the param file and which don't")
	cmd.Flags().Bool("no-keep_state_after_build", false, "If false, Blaze will discard the inmemory state from this build when the build finishes")
	cmd.Flags().Bool("no-track_incremental_state", false, "If false, Blaze will not persist data that allows for invalidation and re- evaluation on incremental builds in order to save memory on this build")
	cmd.Flags().String("skyframe_high_water_mark_full_gc_drops_per_invocation", "", "Flag for advanced configuration of Bazel's internal Skyframe engine")
	cmd.Flags().String("skyframe_high_water_mark_minor_gc_drops_per_invocation", "", "Flag for advanced configuration of Bazel's internal Skyframe engine")
	cmd.Flags().String("skyframe_high_water_mark_threshold", "", "Flag for advanced configuration of Bazel's internal Skyframe engine")
	cmd.Flags().Bool("track_incremental_state", false, "If false, Blaze will not persist data that allows for invalidation and re- evaluation on incremental builds in order to save memory on this build")

	// TODO add flag completion
}
