package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show current preference values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().Bool("json", false, "output as JSON")
	getCmd.Flags().Bool("set-flags", false, "output as \"tailscale set\" flag arguments")
	rootCmd.AddCommand(getCmd)
}
