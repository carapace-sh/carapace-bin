package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showExternalsCmd = &cobra.Command{
	Use:   "show-externals",
	Short: "Show svn:externals listings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showExternalsCmd).Standalone()

	showExternalsCmd.Flags().IntP("revision", "r", 0, "Refer to a specific revision")
	rootCmd.AddCommand(showExternalsCmd)
}
