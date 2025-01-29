package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildx_pruneCmd = &cobra.Command{
	Use:   "prune [options]",
	Short: "Remove unused images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildx_pruneCmd).Standalone()

	buildx_pruneCmd.Flags().BoolP("all", "a", false, "Remove all images not in use by containers, not just dangling ones")
	buildx_pruneCmd.Flags().Bool("build-cache", false, "Remove persistent build cache created by --mount=type=cache")
	buildx_pruneCmd.Flags().Bool("external", false, "Remove images even when they are used by external containers (e.g., by build containers)")
	buildx_pruneCmd.Flags().StringSlice("filter", []string{}, "Provide filter values (e.g. 'label=<key>=<value>')")
	buildx_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	buildxCmd.AddCommand(buildx_pruneCmd)
}
