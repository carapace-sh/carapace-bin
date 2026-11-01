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

	rootCmd.Flags().BoolP("help", "h", false, "print the command line options for eopkg and exit")
	rootCmd.Flags().Bool("version", false, "print the eopkg version and exit")

	rootCmd.PersistentFlags().StringP("bandwidth-limit", "L", "", "keep bandwidth usage under specified KB's")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "show debugging information")
	rootCmd.PersistentFlags().StringP("destdir", "D", "", "change the system root for eopkg commands")
	rootCmd.PersistentFlags().IntP("download-workers", "w", 8, "set the max number of concurrent download workers")
	rootCmd.PersistentFlags().BoolP("no-color", "N", false, "suppresses all coloring of eopkg's output")
	rootCmd.PersistentFlags().StringP("password", "p", "", "set password used when connecting to Basic-Auth repositories")
	rootCmd.PersistentFlags().IntP("retry-attempts", "R", 0, "set the max number of retry attempts in case of connection timeouts")
	rootCmd.PersistentFlags().StringP("username", "u", "", "set username used when connecting to Basic-Auth repositories")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "detailed output")
	rootCmd.PersistentFlags().BoolP("yes-all", "y", false, "assume yes in all yes/no queries")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"destdir": carapace.ActionDirectories(),
	})
}
