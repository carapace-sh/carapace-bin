package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print OpenSSH configuration details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().StringP("port", "p", "", "SSH port on a remote host.")
	rootCmd.AddCommand(configCmd)
}
