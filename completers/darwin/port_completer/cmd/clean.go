package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean files used for building a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()
	cleanCmd.Flags().Bool("all", false, "Remove all build files (dist, archive, logs, work)")
	cleanCmd.Flags().Bool("archive", false, "Remove temporary archives")
	cleanCmd.Flags().Bool("dist", false, "Remove downloaded files")
	cleanCmd.Flags().Bool("logs", false, "Remove log files")
	cleanCmd.Flags().Bool("work", false, "Remove the work directory (default)")
	rootCmd.AddCommand(cleanCmd)
}
