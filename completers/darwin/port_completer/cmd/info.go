package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display meta-information about a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	infoCmd.Flags().Bool("category", false, "Display category")
	infoCmd.Flags().Bool("depends", false, "Display dependencies")
	infoCmd.Flags().Bool("fullname", false, "Display full name with version")
	infoCmd.Flags().Bool("homepage", false, "Display homepage")
	infoCmd.Flags().Bool("index", false, "Use cached port index")
	infoCmd.Flags().Bool("line", false, "Output in single-line format")
	infoCmd.Flags().Bool("maintainer", false, "Display maintainer")
	infoCmd.Flags().Bool("pretty", false, "Pretty-print output")
	rootCmd.AddCommand(infoCmd)
}
