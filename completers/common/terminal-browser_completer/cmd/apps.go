package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Lists registered terminal-browser apps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsCmd).Standalone()

	appsCmd.Flags().Bool("json", false, "Machine readable output")
	rootCmd.AddCommand(appsCmd)
}
