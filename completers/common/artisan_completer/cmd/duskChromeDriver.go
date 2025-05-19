package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duskChromeDriverCmd = &cobra.Command{
	Use:   "dusk:chrome-driver [--all] [--detect] [--proxy [PROXY]] [--ssl-no-verify] [--] [<version>]",
	Short: "Install the ChromeDriver binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duskChromeDriverCmd).Standalone()

	duskChromeDriverCmd.Flags().Bool("all", false, "Install a ChromeDriver binary for every OS")
	duskChromeDriverCmd.Flags().Bool("detect", false, "Detect the installed Chrome / Chromium version")
	duskChromeDriverCmd.Flags().String("proxy", "", "The proxy to download the binary through (example: \"tcp://127.0.0.1:9000\")")
	duskChromeDriverCmd.Flags().Bool("ssl-no-verify", false, "Bypass SSL certificate verification when installing through a proxy")
	rootCmd.AddCommand(duskChromeDriverCmd)
}
