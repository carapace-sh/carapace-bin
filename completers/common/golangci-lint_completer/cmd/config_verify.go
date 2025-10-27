package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify configuration against JSON schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_verifyCmd).Standalone()

	config_verifyCmd.Flags().String("schema", "", "JSON schema URL")
	config_verifyCmd.Flag("schema").Hidden = true
	configCmd.AddCommand(config_verifyCmd)
}
