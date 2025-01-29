package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_pruneCmd = &cobra.Command{
	Use:   "prune [options]",
	Short: "Remove unused data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_pruneCmd).Standalone()

	system_pruneCmd.Flags().BoolP("all", "a", false, "Remove all unused data")
	system_pruneCmd.Flags().Bool("build", false, "Remove build containers")
	system_pruneCmd.Flags().Bool("external", false, "Remove container data in storage not controlled by podman")
	system_pruneCmd.Flags().StringSlice("filter", []string{}, "Provide filter values (e.g. 'label=<key>=<value>')")
	system_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation.  The default is false")
	system_pruneCmd.Flags().Bool("volumes", false, "Prune volumes")
	systemCmd.AddCommand(system_pruneCmd)
}
