package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_userCmd = &cobra.Command{
	Use:   "user",
	Short: "View and configure user information (name, email, editor)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_userCmd).Standalone()

	config_userCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_userCmd)
}
