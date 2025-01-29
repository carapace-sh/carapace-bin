package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_listCmd = &cobra.Command{
	Use:     "list [options] [IMAGE]",
	Short:   "List images in local storage",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_listCmd).Standalone()

	image_listCmd.Flags().BoolP("all", "a", false, "Show all images (default hides intermediate images)")
	image_listCmd.Flags().Bool("digests", false, "Show digests")
	image_listCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions provided (default [])")
	image_listCmd.Flags().String("format", "", "Change the output format to JSON or a Go template")
	image_listCmd.Flags().Bool("history", false, "Display the image name history")
	image_listCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	image_listCmd.Flags().BoolP("noheading", "n", false, "Do not print column headings")
	image_listCmd.Flags().BoolP("quiet", "q", false, "Display only image IDs")
	image_listCmd.Flags().String("sort", "", "Sort by created, id, repository, size, tag")
	imageCmd.AddCommand(image_listCmd)
}
