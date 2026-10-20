package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Inspect bootstrap secret inputs without revealing their values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_secretsCmd).Standalone()

	bootstrap_secretsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_secretsCmd)
}
