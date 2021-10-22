package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_cdn_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a CDN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_cdn_deleteCmd).Standalone()
	compute_cdn_deleteCmd.Flags().BoolP("force", "f", false, "Delete the specified CDN without prompting for confirmation")
	compute_cdnCmd.AddCommand(compute_cdn_deleteCmd)
}
