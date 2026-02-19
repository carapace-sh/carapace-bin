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

	uploadCmd.Flags().String("allow-insecure-host", "", "Hosts that are allowed to use insecure connections")
	uploadCmd.Flags().String("auth-file", "", "Path to the file containing the authentication token")
	rootCmd.AddCommand(uploadCmd)

	carapace.Gen(uploadCmd).FlagCompletion(carapace.ActionMap{
		"auth-file": carapace.ActionFiles(),
	})
}
