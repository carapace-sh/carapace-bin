package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildx_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print build version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildx_versionCmd).Standalone()

	buildxCmd.AddCommand(buildx_versionCmd)
}
