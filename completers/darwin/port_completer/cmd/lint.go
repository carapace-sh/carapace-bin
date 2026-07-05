package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Verify a Portfile conforms to MacPorts standards",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()
	lintCmd.Flags().Bool("nitpick", false, "Enable nitpick mode for stricter checks")
	rootCmd.AddCommand(lintCmd)
}
