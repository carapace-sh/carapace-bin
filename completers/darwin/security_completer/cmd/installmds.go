package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installMdsCmd = &cobra.Command{
	Use:   "install-mds",
	Short: "Install or re-install the MDS database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installMdsCmd).Standalone()
	rootCmd.AddCommand(installMdsCmd)
}
