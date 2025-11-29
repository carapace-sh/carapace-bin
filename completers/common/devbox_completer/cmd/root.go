package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devbox",
	Short: "Instant, easy, predictable development environments",
	Long:  "https://www.jetpack.io/devbox/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("debug", false, "Show full stack traces on errors")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "suppresses logs")
	rootCmd.PersistentFlags().String("trace", "", "write a trace to a file")
	rootCmd.Flag("debug").Hidden = true
	rootCmd.Flag("trace").NoOptDefVal = " "
	rootCmd.Flag("trace").Hidden = true
}
