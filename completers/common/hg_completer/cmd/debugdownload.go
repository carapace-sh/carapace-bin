package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugdownloadCmd = &cobra.Command{
	Use:    "debugdownload",
	Short:  "download a resource using Mercurial logic and config",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugdownloadCmd).Standalone()

	debugdownloadCmd.Flags().StringP("output", "o", "", "path")
	rootCmd.AddCommand(debugdownloadCmd)
}
