package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:     "pull [OPTIONS] NAME[:TAG|@DIGEST]",
	Short:   "Download an image from a registry",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().BoolP("all-tags", "a", false, "Download all tagged images in the repository")
	pullCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (deprecated)")
	pullCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	pullCmd.Flags().BoolP("quiet", "q", false, "Suppress verbose output")
	pullCmd.Flag("disable-content-trust").Hidden = true
	rootCmd.AddCommand(pullCmd)
}
