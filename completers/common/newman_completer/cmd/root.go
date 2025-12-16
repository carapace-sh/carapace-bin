package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "newman",
	Short: "Newman is a command-line collection runner for Postman",
	Long:  "https://github.com/postmanlabs/newman",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display help for command")
	rootCmd.Flags().BoolP("version", "v", false, "output the version number")
}
