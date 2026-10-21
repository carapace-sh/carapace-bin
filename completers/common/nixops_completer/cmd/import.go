package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import deployments into the state file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().Bool("include-keys", false, "Import public SSH host keys to .ssh/known_hosts")

	rootCmd.AddCommand(importCmd)
}
