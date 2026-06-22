package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devcontainerCmd = &cobra.Command{
	Use:   "devcontainer",
	Short: "Generate Dev Container setup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devcontainerCmd).Standalone()
	devcontainerCmd.Flags().BoolP("help", "h", false, "Show help")
}
