package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installsCmd = &cobra.Command{
	Use:   "installs",
	Short: "[experimental] Inspect and choose identity-layout installations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installsCmd).Standalone()

	installsCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(installsCmd)
}
