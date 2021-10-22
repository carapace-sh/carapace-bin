package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_image_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create custom image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_image_createCmd).Standalone()
	compute_image_createCmd.Flags().String("image-description", "", "Description of image")
	compute_image_createCmd.Flags().String("image-distribution", "Unknown", "Custom image distribution")
	compute_image_createCmd.Flags().String("image-url", "", "Custom image retrieval URL (required)")
	compute_image_createCmd.Flags().String("region", "", "Region slug identifier (required)")
	compute_image_createCmd.Flags().StringSlice("tag-names", []string{}, "List of tags applied to image")
	compute_imageCmd.AddCommand(compute_image_createCmd)

	carapace.Gen(compute_image_createCmd).FlagCompletion(carapace.ActionMap{
		"region": action.ActionRegions(),
	})
}
