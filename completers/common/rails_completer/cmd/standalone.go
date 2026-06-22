package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var standaloneCmd = &cobra.Command{
	Use:   "standalone",
	Short: "Standalone Rails commands (about, version, etc.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(standaloneCmd).Standalone()

	standaloneCmd.Flags().BoolP("help", "h", false, "Show help")
}
