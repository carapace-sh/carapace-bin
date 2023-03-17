package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var dlxCmd = &cobra.Command{
	Use:     "dlx",
	Short:   "Run a package in a temporary environment",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlxCmd).Standalone()

	dlxCmd.Flags().StringSliceP("package", "p", []string{}, "The package(s) to install before running the command")
	dlxCmd.Flags().BoolP("quiet", "q", false, "Only report critical errors instead of printing the full install logs")
	rootCmd.AddCommand(dlxCmd)

	// TODO package completion

	embed.EmbedCarapaceBin(dlxCmd)
}
