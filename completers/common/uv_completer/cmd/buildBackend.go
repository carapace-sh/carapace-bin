package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildBackendCmd = &cobra.Command{
	Use:    "build-backend",
	Short:  "The implementation of the build backend",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildBackendCmd).Standalone()

	rootCmd.AddCommand(buildBackendCmd)
}
