package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info [OPTIONS]",
	Short:   "Display system-wide information",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	rootCmd.AddCommand(infoCmd)
}
