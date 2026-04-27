package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eopkg",
	Short: "Solus package manager",
	Long:  "https://help.getsol.us/docs/user/package-management/basics/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("destdir", "D", "", "Change the system root for eopkg run operations")
	rootCmd.PersistentFlags().BoolP("yes-all", "y", false, "Assume yes for all yes/no queries")
	rootCmd.PersistentFlags().StringP("username", "u", "", "Set username")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Set password")
	rootCmd.PersistentFlags().StringP("bandwidth-limit", "L", "", "Keep bandwidth usage under specified KB/s")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Detailed output")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Show debug output")
	rootCmd.PersistentFlags().Bool("no-color", false, "Suppress all coloring of output")
	rootCmd.PersistentFlags().IntP("retry-attempts", "r", 0, "Set the max number of retry attempts in case of connection timeouts")
	rootCmd.PersistentFlags().Bool("version", false, "Show program version")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show help message")
}
