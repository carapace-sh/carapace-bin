package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pacman-db-upgrade",
	Short: "Upgrade the local pacman database to a newer format",
	Long:  "https://manpages.debian.org/testing/pacman-package-manager/pacman-db-upgrade.8.en.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("config", "", "set an alternate configuration file")
	rootCmd.Flags().StringP("dbpath", "d", "", "set an alternate database location")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("nocolor", false, "disable colorized output messages")
	rootCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	rootCmd.Flags().BoolP("version", "V", false, "show version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"dbpath": carapace.ActionDirectories(),
		"root":   carapace.ActionDirectories(),
	})
}
