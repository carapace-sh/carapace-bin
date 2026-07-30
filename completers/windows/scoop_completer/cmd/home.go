package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "open the package homepage in a browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeCmd).Standalone()
	rootCmd.AddCommand(homeCmd)
}
