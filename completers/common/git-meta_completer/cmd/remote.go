package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage metadata remote sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteCmd).Standalone()

	remoteCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(remoteCmd)
}
