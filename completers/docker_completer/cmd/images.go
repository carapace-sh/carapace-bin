package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:     "images [OPTIONS] [REPOSITORY[:TAG]]",
	Short:   "List images",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagesCmd).Standalone()

	imagesCmd.Flags().BoolP("all", "a", false, "Show all images (default hides intermediate images)")
	imagesCmd.Flags().Bool("digests", false, "Show digests")
	imagesCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	imagesCmd.Flags().String("format", "", "Format output using a custom template:")
	imagesCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	imagesCmd.Flags().BoolP("quiet", "q", false, "Only show image IDs")
	rootCmd.AddCommand(imagesCmd)

	carapace.Gen(imagesCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
