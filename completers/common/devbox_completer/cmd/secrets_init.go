package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secrets_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize secrets management with jetify cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_initCmd).Standalone()

	secrets_initCmd.Flags().BoolP("force", "f", false, "Force initialization even if already initialized")
	secretsCmd.AddCommand(secrets_initCmd)
}
