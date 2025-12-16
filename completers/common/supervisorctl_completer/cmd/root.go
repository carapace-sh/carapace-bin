package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "supervisorctl",
	Short: "control applications run by supervisord from the cmd line",
	Long:  "http://supervisord.org/",
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

	rootCmd.Flags().StringP("configuration", "c", "", "configuration file path")
	rootCmd.Flags().BoolP("help", "h", false, "print usage message and exit")
	rootCmd.Flags().BoolP("history-file", "r", false, "keep a readline history")
	rootCmd.Flags().BoolP("interactive", "i", false, "start an interactive shell after executing commands")
	rootCmd.Flags().StringP("password", "p", "", "password to use for authentication with server")
	rootCmd.Flags().StringP("serverurl", "s", "", "URL on which supervisord server is listening")
	rootCmd.Flags().StringP("username", "u", "", "username to use for authentication with server")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionFiles(),
	})
}
