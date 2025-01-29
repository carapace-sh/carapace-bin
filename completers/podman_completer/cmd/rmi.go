package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rmiCmd = &cobra.Command{
	Use:   "rmi [options] IMAGE [IMAGE...]",
	Short: "Remove one or more images from local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmiCmd).Standalone()

	rmiCmd.Flags().BoolP("all", "a", false, "Remove all images")
	rmiCmd.Flags().BoolP("force", "f", false, "Force Removal of the image")
	rmiCmd.Flags().BoolP("ignore", "i", false, "Ignore errors if a specified image does not exist")
	rmiCmd.Flags().Bool("no-prune", false, "Do not remove dangling images")
	rootCmd.AddCommand(rmiCmd)
}
