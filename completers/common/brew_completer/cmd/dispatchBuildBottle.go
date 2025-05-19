package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatchBuildBottleCmd = &cobra.Command{
	Use:     "dispatch-build-bottle",
	Short:   "Build bottles for these formulae with GitHub Actions",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatchBuildBottleCmd).Standalone()

	dispatchBuildBottleCmd.Flags().Bool("debug", false, "Display any debugging information.")
	dispatchBuildBottleCmd.Flags().Bool("help", false, "Show this message.")
	dispatchBuildBottleCmd.Flags().Bool("issue", false, "If specified, post a comment to this issue number if the job fails.")
	dispatchBuildBottleCmd.Flags().Bool("linux", false, "Dispatch bottle for Linux (using GitHub runners).")
	dispatchBuildBottleCmd.Flags().Bool("linux-self-hosted", false, "Dispatch bottle for Linux (using self-hosted runner).")
	dispatchBuildBottleCmd.Flags().Bool("linux-wheezy", false, "Use Debian Wheezy container for building the bottle on Linux.")
	dispatchBuildBottleCmd.Flags().Bool("macos", false, "macOS version (or comma-separated list of versions) the bottle should be built for.")
	dispatchBuildBottleCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	dispatchBuildBottleCmd.Flags().Bool("tap", false, "Target tap repository (default: `homebrew/core`).")
	dispatchBuildBottleCmd.Flags().Bool("timeout", false, "Build timeout (in minutes, default: 60).")
	dispatchBuildBottleCmd.Flags().Bool("upload", false, "Upload built bottles.")
	dispatchBuildBottleCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	dispatchBuildBottleCmd.Flags().Bool("workflow", false, "Dispatch specified workflow (default: `dispatch-build-bottle.yml`).")
	rootCmd.AddCommand(dispatchBuildBottleCmd)
}
