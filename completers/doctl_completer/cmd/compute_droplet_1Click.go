package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_1ClickCmd = &cobra.Command{
	Use:   "1-click",
	Short: "Display commands that pertain to Droplet 1-click applications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_1ClickCmd).Standalone()
	compute_dropletCmd.AddCommand(compute_droplet_1ClickCmd)
}
