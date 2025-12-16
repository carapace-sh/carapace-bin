package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apt",
	Short: "apt is a commandline package manager",
	Long:  "https://salsa.debian.org/apt-team/apt",
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

	rootCmd.Flags().StringS("c", "c", "", "config file")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().StringP("host-architecture", "a", "", "architecture")
	rootCmd.Flags().StringP("option", "o", "", "config string")
	rootCmd.Flags().StringP("target-release", "t", "", "target release")
	rootCmd.Flags().BoolP("version", "v", false, "show version")

	// TODO release/architecture
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionFiles(),
	})
}
