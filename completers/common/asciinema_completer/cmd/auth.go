package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate this CLI with an asciinema server account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(authCmd).Standalone()

	authCmd.Flags().BoolP("help", "h", false, "Print help")
	authCmd.Flags().String("server-url", "", "asciinema server URL")
	rootCmd.AddCommand(authCmd)
}
