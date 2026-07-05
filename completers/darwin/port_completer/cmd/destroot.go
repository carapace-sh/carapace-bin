package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var destrootCmd = &cobra.Command{
	Use:   "destroot",
	Short: "Install a port to a temporary directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destrootCmd).Standalone()
	rootCmd.AddCommand(destrootCmd)
}
