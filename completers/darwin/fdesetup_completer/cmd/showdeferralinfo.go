package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showdeferralinfoCmd = &cobra.Command{
	Use:   "showdeferralinfo",
	Short: "Display deferral configuration information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showdeferralinfoCmd).Standalone()
	rootCmd.AddCommand(showdeferralinfoCmd)
}
