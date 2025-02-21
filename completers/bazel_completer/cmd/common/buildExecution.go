package common

import (
	"github.com/spf13/cobra"
)

func AddBuildExecutionFlags(cmd *cobra.Command) {
	cmd.Flags().String("experimental_ui_max_stdouterr_bytes", "", "The maximum size of the stdout / stderr files that will be printed to the console")
	cmd.Flags().String("gc_thrashing_threshold", "", "The percent of tenured space occupied (0-100) above which GcThrashingDetector considers memory pressure events against its limits (-- gc_thrashing_limits)")
	cmd.Flags().Bool("incompatible_remote_dangling_symlinks", false, "If set to true, symlinks uploaded to a remote or disk cache are allowed to dangle")
	cmd.Flags().Bool("incompatible_remote_symlinks", false, "If set to true, Bazel will always upload symlinks as such to a remote or disk cache")
	cmd.Flags().Bool("no-incompatible_remote_dangling_symlinks", false, "If set to true, symlinks uploaded to a remote or disk cache are allowed to dangle")
	cmd.Flags().Bool("no-incompatible_remote_symlinks", false, "If set to true, Bazel will always upload symlinks as such to a remote or disk cache")
}
