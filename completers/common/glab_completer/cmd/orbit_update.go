package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orbit_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the GitLab Orbit CLI binary to the latest version. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_updateCmd).Standalone()

	orbit_updateCmd.Flags().BoolP("yes", "y", false, "Skip the download prompt when the binary is not installed.")
	orbitCmd.AddCommand(orbit_updateCmd)
}
