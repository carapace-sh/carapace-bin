package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show the state of every configured bootstrap part",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_statusCmd).Standalone()

	bootstrap_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if anything differs from the config (all entries are still listed)")
	bootstrap_statusCmd.Flags().Bool("prompt-secrets", false, "Prompt securely for missing bootstrap secret inputs")
	bootstrapCmd.AddCommand(bootstrap_statusCmd)
}
