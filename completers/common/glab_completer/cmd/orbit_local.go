package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orbit_localCmd = &cobra.Command{
	Use:   "local [command]",
	Short: "Run the Orbit local CLI (Experimental)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbit_localCmd).Standalone()

	orbit_localCmd.Flags().Bool("install", false, "Install the Orbit local CLI binary without running it.")
	orbit_localCmd.Flags().Bool("update", false, "Check for and install updates to the binary.")
	orbit_localCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts.")
	orbitCmd.AddCommand(orbit_localCmd)
}
