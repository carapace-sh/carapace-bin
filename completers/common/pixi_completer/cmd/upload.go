package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload conda packages to various channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploadCmd).Standalone()

	uploadCmd.Flags().StringSlice("allow-insecure-host", nil, "List of hosts for which SSL certificate verification should be skipped")
	rootCmd.AddCommand(uploadCmd)
}
