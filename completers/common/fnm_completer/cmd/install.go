package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a new Node.js version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("latest", false, "Install latest version")
	installCmd.Flags().Bool("lts", false, "Install latest LTS")
	installCmd.Flags().String("progress", "auto", "Show an interactive progress bar for the download status")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"progress": carapace.ActionValues("auto", "never", "always"),
	})
}
