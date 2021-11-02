package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_firewall_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Add tags to a cloud firewall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_firewall_addTagsCmd).Standalone()
	compute_firewall_addTagsCmd.Flags().StringSlice("tag-names", []string{}, "A comma-separated list of tag names to apply to the cloud firewall, e.g.: `frontend,backend`")
	compute_firewallCmd.AddCommand(compute_firewall_addTagsCmd)
}
