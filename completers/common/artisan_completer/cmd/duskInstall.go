package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duskInstallCmd = &cobra.Command{
	Use:   "dusk:install [--proxy [PROXY]] [--ssl-no-verify]",
	Short: "Install Dusk into the application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duskInstallCmd).Standalone()

	duskInstallCmd.Flags().String("proxy", "", "The proxy to download the binary through (example: \"tcp://127.0.0.1:9000\")")
	duskInstallCmd.Flags().Bool("ssl-no-verify", false, "Bypass SSL certificate verification when installing through a proxy")
	rootCmd.AddCommand(duskInstallCmd)
}
