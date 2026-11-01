package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_forgejoCmd = &cobra.Command{
	Use:   "forgejo",
	Short: "Show the Forgejo token mise uses for a host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_forgejoCmd).Standalone()

	token_forgejoCmd.Flags().BoolP("help", "h", false, "Print help")
	token_forgejoCmd.Flags().Bool("unmask", false, "Show the full unmasked token")
	tokenCmd.AddCommand(token_forgejoCmd)
}
