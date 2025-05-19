package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down [--redirect [REDIRECT]] [--render [RENDER]] [--retry [RETRY]] [--refresh [REFRESH]] [--secret [SECRET]] [--with-secret] [--status [STATUS]]",
	Short: "Put the application into maintenance / demo mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downCmd).Standalone()

	downCmd.Flags().String("redirect", "", "The path that users should be redirected to")
	downCmd.Flags().String("refresh", "", "The number of seconds after which the browser may refresh")
	downCmd.Flags().String("render", "", "The view that should be prerendered for display during maintenance mode")
	downCmd.Flags().String("retry", "", "The number of seconds after which the request may be retried")
	downCmd.Flags().String("secret", "", "The secret phrase that may be used to bypass maintenance mode")
	downCmd.Flags().String("status", "", "The status code that should be used when returning the maintenance mode response")
	downCmd.Flags().Bool("with-secret", false, "Generate a random secret phrase that may be used to bypass maintenance mode")
	rootCmd.AddCommand(downCmd)
}
