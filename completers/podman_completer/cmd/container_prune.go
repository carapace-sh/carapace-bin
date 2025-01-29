package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_pruneCmd = &cobra.Command{
	Use:   "prune [options]",
	Short: "Remove all non running containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_pruneCmd).Standalone()

	container_pruneCmd.Flags().StringSlice("filter", []string{}, "Provide filter values (e.g. 'label=<key>=<value>')")
	container_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation.  The default is false")
	containerCmd.AddCommand(container_pruneCmd)
}
