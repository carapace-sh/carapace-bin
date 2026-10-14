package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "exports installed apps, buckets and optionally configs in JSON format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()
	exportCmd.Flags().BoolP("config", "c", false, "export the Scoop configuration file too")
	rootCmd.AddCommand(exportCmd)
}
