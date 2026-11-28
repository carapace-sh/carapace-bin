package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open web browser with documentation for ESP-IDF",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsCmd).Standalone()

	docsCmd.Flags().StringP("language", "l", "", "Documentation language")
	docsCmd.Flags().Bool("no-browser", false, "Don't open browser")
	docsCmd.Flags().String("starting-page", "", "Documentation page")
	docsCmd.Flags().StringP("target", "t", "", "Chip target")
	docsCmd.Flags().String("version", "", "Version of ESP-IDF")
	rootCmd.AddCommand(docsCmd)

	carapace.Gen(docsCmd).FlagCompletion(carapace.ActionMap{
		"language": carapace.ActionValues("cn", "en", "zh_CN"),
		"target":   espidf.ActionTargets(),
	})
}
