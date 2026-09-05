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

	rootCmd.PersistentFlags().StringP("bandwidth-limit", "L", "", "keep bandwidth usage under the specified (numeric) KBs")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable full debug information and backtraces")
	rootCmd.PersistentFlags().StringP("destdir", "D", "", "change the system root for eopkg commands")
	rootCmd.PersistentFlags().BoolP("no-color", "N", false, "disable the use of ANSI escape sequences for colourisation")
	rootCmd.PersistentFlags().StringP("password", "p", "", "set password used when connecting to Basic-Auth repositories")
	rootCmd.PersistentFlags().Int("retry-attempts", 0, "number of retry attempts for operations that fail")
	rootCmd.PersistentFlags().StringP("username", "u", "", "set username used when connecting to Basic-Auth repositories")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "detailed output")
	rootCmd.PersistentFlags().BoolP("yes-all", "y", false, "assume yes in all yes/no queries")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"destdir": carapace.ActionDirectories(),
	})
}
